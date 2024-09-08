package infrastructure

import (
	"fmt"
	"log"

	"github.com/andrefelizardo/todo-api/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.Name, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database", err)
		return nil, err
	}

	return db, nil
}