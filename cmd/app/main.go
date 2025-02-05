package main

import (
	"log"

	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
