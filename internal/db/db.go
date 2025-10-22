package db

import (
	"log"

	"github.com/robertbonadeo/go-bootcamp-project/config"
	"github.com/robertbonadeo/go-bootcamp-project/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.GetDBConnString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Content{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
