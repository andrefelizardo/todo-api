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

func TestUserUseCase_CreateUser_InvalidInput(t *testing.T) {
	// mockRepo := new(MockUserRepository)
	usecase := NewUserUseCase(nil)

	testCases := []struct {
		name     string
		input    request.CreateUserRequest
		expectedErr string
	}{
		{
			name: "Name too short",
			input: request.CreateUserRequest{
				Name:     "Jo",
				Email:    "test@example.com",
				Password: "password",
			},
			expectedErr: "Field 'Name' must be at least 3 characters long",
		},
		{
			name: "Email invalid",
			input: request.CreateUserRequest{
				Name:     "John Doe",
				Email:    "test",
				Password: "password",
			},
			expectedErr: "Field 'Email' must be a valid email address",
		},
		{
			name: "Password too short",
			input: request.CreateUserRequest{
				Name:     "John Doe",
				Email:    "test@email.com",
				Password: "pass",
			},
			expectedErr: "Field 'Password' must be at least 6 characters long",
		},
		{
			name: "Name is required",
			input: request.CreateUserRequest{
				Name:     "",
				Email:    "test@email.com",
				Password: "password",
			},
			expectedErr: "Field 'Name' is required",
		},
		{
			name: "Email is required",
			input: request.CreateUserRequest{
				Name:     "John Doe",
				Email:    "",
				Password: "password",
			},
			expectedErr: "Field 'Email' is required",
		},
		{
			name: "Password is required",
			input: request.CreateUserRequest{
				Name:     "John Doe",
				Email:    "test@email.com",
				Password: "",
			},
			expectedErr: "Field 'Password' is required",
		},
		{
				name: "Name is too long",
				input: request.CreateUserRequest{
					Name:     "Aiuhsdfiuahsfiuda iausdhfiuasdhfiuasdh iausdhfiashdfauisdh",
					Email:    "test@email.com",
					Password: "password",
				},
				expectedErr: "Field 'Name' must be at most 50 characters long",
		},
		{
			name: "Password is too long",
			input: request.CreateUserRequest{
				Name:     "John Doe",
				Email:    "test@email.com",
				Password: "Aiuhsdfiuahsfiuda iausdhfiuasdhfiuasdh iausdhfiashdfauisdh",
			},
			expectedErr: "Field 'Password' must be at most 50 characters long",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userResponse, err := usecase.CreateUser(tc.input)

			assert.Error(t, err)
			assert.Nil(t, userResponse)
			assert.EqualError(t, err,  tc.expectedErr)
	})}
}