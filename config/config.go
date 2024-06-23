package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App
		HTTP
		Database
		JWT
	}

	// App -.
	App struct {
		Name    string `env-required:"true" env:"APP_NAME"`
		Version string `env-required:"true" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
	}

	// Database -.
	Database struct {
		URL string `env-required:"true" env:"DATABASE_URL"`
	}

	// JWT -.
	JWT struct {
		Secret          string `env-required:"true" env:"JWT_SECRET"`
		ExpireInMinutes int    `env-required:"true" env:"JWT_EXPIRE_IN_MINUTES"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Config error : ", err)
	}

	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
