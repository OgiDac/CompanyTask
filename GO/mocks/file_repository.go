package mocks

import (
	"context"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/stretchr/testify/mock"
)

type FileRepository struct {
	mock.Mock
}

func (m *FileRepository) SaveUserFile(ctx context.Context, file *domain.UserFile) error {
	args := m.Called(ctx, file)
	return args.Error(0)
}

func (m *FileRepository) GetFileByID(ctx context.Context, id string) (*domain.UserFile, error) {
	args := m.Called(ctx, id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.UserFile), args.Error(1)
}

func (m *FileRepository) DeleteFilesByUserID(ctx context.Context, userID uint) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *FileRepository) GetFilesByUserID(ctx context.Context, userID uint) ([]*domain.UserFile, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*domain.UserFile), args.Error(1)
}
