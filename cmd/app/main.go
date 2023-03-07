package main

import (
	"context"
	"go-values-generator/internal/app"
	"go-values-generator/internal/config"
	"log"
)

// TODO: Docker
// TODO: Unit-Tests
// TODO: ReadME file
// TODO: Swagger documentation

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("config initializing")
	cfg := config.GetConfig()

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(ctx, err)
	}

	log.Println("running application")
	if err := a.Run(ctx); err != nil {
		return
	}
}
