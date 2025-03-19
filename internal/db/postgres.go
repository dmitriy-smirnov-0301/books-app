package db

import (
	"books-app/internal/config"
	"books-app/internal/domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Connect() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if !database.Migrator().HasTable(&domain.Book{}) {
		log.Println("Table 'books' does not exist. Running AutoMigrate()...")
		if err := database.AutoMigrate(&domain.Book{}); err != nil {
			return err
		}
	} else {
		log.Println("Table 'books' already exists. Skipping migration.")
	}

	log.Println("Connected to database successfully.")

	db.connection = database

	return nil
}

func (db *Database) GetConnection() *gorm.DB {
	return db.connection
}
