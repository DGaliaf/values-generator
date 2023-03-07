package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	HTTP struct {
		IP   string `yaml:"ip" env:"HTTP_IP" env-default:"localhost"`
		Port int    `yaml:"port" env:"HTTP_PORT" env-default:"9000"`
	} `yaml:"http"`
	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST" env-default:"0.0.0.0"`
		Port     int    `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
		Database int    `yaml:"database" env:"REDIS_DB" env-default:"0"`
	} `yaml:"redis"`
}

const (
	configPath = "configs/config.local.yml"
)

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("config init")

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			log.Println("Cant`t read environment variables from neither .yaml nor .env")
			log.Println(err)

			err := cleanenv.ReadEnv(instance)
			if err != nil {
				help, _ := cleanenv.GetDescription(instance, nil)
				log.Println(help)
				log.Fatalln(err)
			}
		}
	})
	return instance
}
