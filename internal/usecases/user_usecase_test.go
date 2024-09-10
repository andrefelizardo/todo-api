package usecases

import (
	"testing"

	"github.com/andrefelizardo/todo-api/internal/domain"
	"github.com/andrefelizardo/todo-api/internal/request"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUseCase(mockRepo)

	input := request.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	expectedUser := &domain.User{
		ID:        uuid.New(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  "hashed_password",
	}

	mockRepo.On("Create", mock.AnythingOfType("domain.User")).Return(expectedUser, nil)

	userResponse, err := usecase.CreateUser(input)

	assert.NoError(t, err)
	assert.NotNil(t, userResponse)
	assert.Equal(t, "John Doe", userResponse.Name)
	assert.Equal(t, "john@example.com", userResponse.Email)

	mockRepo.AssertExpectations(t)
}