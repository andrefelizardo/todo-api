package usecases

import (
	"fmt"

	"github.com/andrefelizardo/todo-api/internal/domain"
	"github.com/andrefelizardo/todo-api/internal/repositories"
	"github.com/andrefelizardo/todo-api/internal/request"
	"github.com/go-playground/validator"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(input request.CreateUserRequest) (*domain.User, error) {
	err := u.validateInput(input)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}

	dbUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return dbUser, nil
	
}

func (u *UserUseCase) validateInput(input request.CreateUserRequest) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return fmt.Errorf("Validation error: %s", errors)
	}

	return nil
}