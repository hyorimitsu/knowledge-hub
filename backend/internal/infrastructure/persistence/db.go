package persistence

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(
		&model.Tenant{},
		&model.Knowledge{},
		&model.Tag{},
		&model.Comment{},
		&model.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return &Database{db}
}

func (db *Database) Begin() *gorm.DB {
	return db.DB.Begin()
}

func (db *Database) Transaction(fc func(tx *gorm.DB) error) error {
	return db.DB.Transaction(fc)
}
