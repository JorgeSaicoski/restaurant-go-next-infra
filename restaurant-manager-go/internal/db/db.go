package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database struct encapsulates the database connection.
type Database struct {
	DB *gorm.DB
}

// NewDatabase initializes the Database struct.
func NewDatabase() *Database {
	return &Database{}
}

// Connect establishes the database connection using environment variables.
func (d *Database) Connect() error {
	// Read database configuration from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	d.DB = db

	// Perform auto-migration (modify with actual models)
	if err := db.AutoMigrate( /* Add your models here */ ); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// GetDB returns the database instance.
func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
