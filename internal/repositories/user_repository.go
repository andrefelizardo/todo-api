package repositories

import (
	"github.com/andrefelizardo/todo-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.User) (*domain.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) Create(user domain.User) (*domain.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}