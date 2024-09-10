package usecases

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/andrefelizardo/todo-api/internal/domain"
	"github.com/andrefelizardo/todo-api/internal/repositories"
	"github.com/andrefelizardo/todo-api/internal/request"
	"github.com/andrefelizardo/todo-api/internal/response"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(input request.CreateUserRequest) (*response.UserResponse, error) {
	err := u.validateInput(input)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		log.Error("Error hashing password", err)
		return nil, err
	}

	user := domain.User{
		Name: input.Name,
		Email: input.Email,
		Password: hashedPassword,
	}

	dbUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID: dbUser.ID.String(),
		Name: dbUser.Name,
		Email: dbUser.Email,
		CreatedAt: dbUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: dbUser.UpdatedAt.Format(time.RFC3339),
	}, nil
	
}

func (u *UserUseCase) validateInput(input request.CreateUserRequest) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return fmt.Errorf("validation error: %s", errors)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}