package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/OgiDac/CompanyTask/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUploadFile_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockFileRepo := new(mocks.FileRepository)

	useCase := NewFileUseCase(mockUserRepo, mockFileRepo, 2*time.Second)

	mockUserRepo.On("GetUserByID", mock.Anything, uint(1)).Return(&domain.User{ID: 1}, nil)
	mockFileRepo.On("SaveUserFile", mock.Anything, mock.Anything).Return(nil)

	err := useCase.UploadFile(context.Background(), 1, "file.txt", "text/plain", []byte("data"))

	require.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
	mockFileRepo.AssertExpectations(t)
}

func TestUploadFile_UserNotFound(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockFileRepo := new(mocks.FileRepository)

	useCase := NewFileUseCase(mockUserRepo, mockFileRepo, 2*time.Second)

	// Correctly simulate user not found
	mockUserRepo.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, errors.New("user not found"))

	err := useCase.UploadFile(context.Background(), 2, "file.txt", "text/plain", []byte("data"))

	require.Error(t, err)
	require.Equal(t, "user not found", err.Error())

	mockUserRepo.AssertExpectations(t)
}

func TestGetFileByID_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockFileRepo := new(mocks.FileRepository)

	useCase := NewFileUseCase(mockUserRepo, mockFileRepo, 2*time.Second)

	expectedFile := &domain.UserFile{
		ID:       "abc123",
		Filename: "file.txt",
	}

	mockFileRepo.On("GetFileByID", mock.Anything, "abc123").Return(expectedFile, nil)

	result, err := useCase.GetFileByID(context.Background(), "abc123")

	require.NoError(t, err)
	require.Equal(t, expectedFile, result)
	mockFileRepo.AssertExpectations(t)
}

func TestGetFileByID_NotFound(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockFileRepo := new(mocks.FileRepository)

	useCase := NewFileUseCase(mockUserRepo, mockFileRepo, 2*time.Second)

	mockFileRepo.On("GetFileByID", mock.Anything, "notfound").Return(nil, errors.New("not found"))

	result, err := useCase.GetFileByID(context.Background(), "notfound")

	require.Error(t, err)
	require.Nil(t, result)
	require.EqualError(t, err, "not found")
	mockFileRepo.AssertExpectations(t)
}
