package infrastructure

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return db
}
