package infrastructure

import (
	"fmt"
	"log"

	"github.com/andrefelizardo/todo-api/internal/configs"
	"github.com/andrefelizardo/todo-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *configs.DBConfig) (*gorm.DB, error) {
	fmt.Printf("Connecting to database %s on %s:%s with %s\n", config.Name, config.Host, config.Port, config.SSLMode)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.Name, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database", err)
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		log.Printf("Error during migration: %v", err)
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}