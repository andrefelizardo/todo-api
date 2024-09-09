package repositories

import (
	"github.com/andrefelizardo/todo-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user domain.User) (*domain.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}