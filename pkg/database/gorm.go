package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect and return *gorm.DB connection
func ConnectGorm(dbURL string) *gorm.DB {
	fmt.Println("[GORM] Database connecting...")
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("[GORM] Database connection error : ", err)
	}

	fmt.Println("[GORM] Database connected successfully")
	return db
}
