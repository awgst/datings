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
		GinMode string `env:"GIN_MODE"`
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

func NewConfigForTest() (*Config, error) {
	cfg := &Config{
		App: App{
			Name:    "dating",
			Version: "1.0.0",
			GinMode: "test",
		},
		HTTP: HTTP{
			Port: "8080",
		},
		Database: Database{
			URL: "root:root@tcp(localhost:3306)/datings?charset=utf8mb4&parseTime=True&loc=Local",
		},
		JWT: JWT{
			Secret:          "secret",
			ExpireInMinutes: 60,
		},
	}

	return cfg, nil
}
