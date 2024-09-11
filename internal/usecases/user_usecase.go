package usecases

import (
	"fmt"
	"strings"
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

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%v", v.Message)
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

	if u.userRepository == nil {
        return nil, fmt.Errorf("user repository is not initialized")
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
		var errorMessages []string
		
		for _, e := range errors {
			var errorMsg string

			switch e.Tag() {
			case "required":
				errorMsg = fmt.Sprintf("Field '%s' is required", e.Field())
			case "min":
				errorMsg = fmt.Sprintf("Field '%s' must be at least %s characters long", e.Field(), e.Param())
			case "max":
				errorMsg = fmt.Sprintf("Field '%s' must be at most %s characters long", e.Field(), e.Param())
			case "email":
				errorMsg = fmt.Sprintf("Field '%s' must be a valid email address", e.Field())
			default:
				errorMsg = fmt.Sprintf("Field '%s' is invalid", e.Field())
		}

		errorMessages = append(errorMessages, errorMsg)
	}

	return &ValidationError{Message: strings.Join(errorMessages, ", ")}
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