package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/imrkaofficial/jwtssoauthapp/models"
	"log"
)

var DB *gorm.DB

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error performing auto migration: %v", err)
	}

	DB = db
	return db
}
