package database

import (
	"gorm.io/gorm"
)

type Connection struct {
	Gorm *gorm.DB
}

type ConnectionConfig struct {
	DbUrl string
}

func New(cfg ConnectionConfig) *Connection {
	return &Connection{
		Gorm: ConnectGorm(cfg.DbUrl),
	}
}
