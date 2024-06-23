package app

import (
	"github.com/awgst/datings/config"
	"github.com/awgst/datings/pkg/database"
	"github.com/awgst/datings/pkg/logger"
)

type App struct {
	Config *config.Config
	DB     *database.Connection
	Logger *logger.Logger
}
