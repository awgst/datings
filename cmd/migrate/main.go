package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awgst/datings/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%s", cfg.Database.URL),
	)
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		help()
	}

	switch args[0] {
	case "up":
		up(m)
	case "down":
		down(m)
	case "status":
		status(m)
	default:
		help()
	}
}

func up(m *migrate.Migrate) {
	fmt.Println("Migrating...")
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func down(m *migrate.Migrate) {
	fmt.Println("Rolling back...")
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func status(m *migrate.Migrate) {
	v, _, err := m.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  up")
	fmt.Println("  down")
	fmt.Println("  status")
	fmt.Println("Example: ./migrate up")

	os.Exit(1)
}
