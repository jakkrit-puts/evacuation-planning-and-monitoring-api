package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	AppPort    string `env:"APP_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBUsername string `env:"DB_USERNAME,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBSSLMode  string `env:"DB_SSL_MODE,required"`
}

func NewEnvConfig() *EnvConfig {
	fmt.Println("-- load .env file --")
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load the .env => %e", err)
	}

	cfg := &EnvConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Unable to load variables from the .env => %e", err)
	}

	return cfg
}
