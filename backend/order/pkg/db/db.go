package db

import (
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func InitDB(cfg *config.Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgreDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Order{})

	return &DB{db}, nil
}
