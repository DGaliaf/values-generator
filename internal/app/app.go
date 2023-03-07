package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-values-generator/internal/adapters/db/redis"
	"go-values-generator/internal/config"
	v1 "go-values-generator/internal/controllers/http/v1"
	"go-values-generator/internal/domain/service"
	r "go-values-generator/pkg/client/redis"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"

	"time"
)

type App struct {
	cfg *config.Config

	router chi.Router
}

func NewApp(cfg *config.Config) (*App, error) {
	log.Println("router initializing")
	router := chi.NewRouter()
	router.Use(middleware.RequestID)

	log.Println("swagger docs initializing")
	router.Mount("/swagger", httpSwagger.WrapHandler)

	log.Println("client initializing")
	redisClient, err := r.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	valueStorage := redis.NewValueStorage(redisClient)
	valueService := service.NewValueService(valueStorage)
	valueHandler := v1.NewValueHandler(valueService)
	valueHandler.Register(router)

	return &App{
		cfg:    cfg,
		router: router,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, _ := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP()
	})

	return grp.Wait()
}

func (a *App) startHTTP() error {
	log.Println("HTTP Server initializing")

	log.Println(fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		log.Fatalln("failed to create listener")
	}

	httpServer := &http.Server{
		Handler:      a.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err = httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("server shutdown")
		default:
			log.Fatal(err)
		}
	}

	if err = httpServer.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

	return err
}
