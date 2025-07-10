package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/OgiDac/CompanyTask/config"
	"github.com/OgiDac/CompanyTask/domain"
	"github.com/OgiDac/CompanyTask/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func getTestEnv() *config.Env {
	return &config.Env{
		AccessTokenSecret:      "testsecret",
		AccessTokenExpiryHour:  2,
		RefreshTokenSecret:     "testsecret",
		RefreshTokenExpiryHour: 168,
	}
}

func TestCreateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockPublisher := &mocks.Publisher{}
	env := getTestEnv()
	useCase := NewUserUseCase(mockUserRepo, mockPublisher, 2*time.Second, env)

	req := domain.SignUpRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepassword",
	}

	mockUserRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	access, refresh, err := useCase.CreateUser(context.Background(), req)

	require.NoError(t, err)
	require.NotEmpty(t, access)
	require.NotEmpty(t, refresh)
	require.Len(t, mockPublisher.Published, 1)
	require.Equal(t, "UserCreated", mockPublisher.Published[0].Type)

	mockUserRepo.AssertExpectations(t)
}

func TestCreateUser_EmailExistsError(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockPublisher := &mocks.Publisher{}
	env := getTestEnv()
	useCase := NewUserUseCase(mockUserRepo, mockPublisher, 2*time.Second, env)

	req := domain.SignUpRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepassword",
	}

	mockUserRepo.On("CreateUser", mock.Anything, mock.Anything).Return(errors.New("email already exists"))

	access, refresh, err := useCase.CreateUser(context.Background(), req)

	require.Error(t, err)
	require.EqualError(t, err, "email already exists")
	require.Empty(t, access)
	require.Empty(t, refresh)
	require.Len(t, mockPublisher.Published, 0)

	mockUserRepo.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockPublisher := &mocks.Publisher{}
	env := getTestEnv()
	useCase := NewUserUseCase(mockUserRepo, mockPublisher, 2*time.Second, env)

	mockUserRepo.On("GetUsers", mock.Anything).Return([]*domain.User{
		{ID: 1, Name: "John", Email: "john@example.com"},
	}, nil)

	users, err := useCase.GetAllUsers(context.Background())

	require.NoError(t, err)
	require.Len(t, users, 1)
	require.Equal(t, uint(1), users[0].ID)
	require.Equal(t, "John", users[0].Name)
	require.Equal(t, "john@example.com", users[0].Email)

	mockUserRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockPublisher := &mocks.Publisher{}
	env := getTestEnv()
	useCase := NewUserUseCase(mockUserRepo, mockPublisher, 2*time.Second, env)

	req := domain.UpdateRequest{
		Id:    1,
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	mockUserRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(nil)

	err := useCase.UpdateUser(context.Background(), req)

	require.NoError(t, err)
	require.Len(t, mockPublisher.Published, 1)
	require.Equal(t, "UserUpdated", mockPublisher.Published[0].Type)

	mockUserRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockPublisher := &mocks.Publisher{}
	env := getTestEnv()
	useCase := NewUserUseCase(mockUserRepo, mockPublisher, 2*time.Second, env)

	mockUserRepo.On("DeleteUser", mock.Anything, uint(1)).Return(nil)

	err := useCase.DeleteUser(context.Background(), 1)

	require.NoError(t, err)
	require.Len(t, mockPublisher.Published, 1)
	require.Equal(t, "UserDeleted", mockPublisher.Published[0].Type)

	mockUserRepo.AssertExpectations(t)
}
