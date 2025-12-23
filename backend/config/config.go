package config

import (
	"errors"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string         `env:"ENV" envDefault:"development"`
	Host     string         `env:"HOST" envDefault:"localhost"`
	Port     string         `env:"PORT" envDefault:"5000"`
	Database DatabaseConfig `envPrefix:"DATABASE_"`
	JWT      JWTConfig      `envPrefix:"JWT_"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"3306"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:""`
	Database string `env:"DATABASE" envDefault:"app_db"`
}
type JWTConfig struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`
	ExpiresAt int    `env:"EXPIRES_AT" envDefault:"1"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("ERROR LOADING .ENV FILE")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("ERROR PARSING ENVIRONMENT VARIABLES")
	}

	return &cfg, nil
}
