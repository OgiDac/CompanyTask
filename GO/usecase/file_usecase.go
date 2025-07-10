package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/OgiDac/CompanyTask/repository"
)

type fileUseCase struct {
	userRepo repository.UserRepository
	fileRepo repository.FileRepository
	timeout  time.Duration
}

func NewFileUseCase(userRepo repository.UserRepository, fileRepo repository.FileRepository, timeout time.Duration) domain.FileUseCase {
	return &fileUseCase{
		userRepo: userRepo,
		fileRepo: fileRepo,
		timeout:  timeout,
	}
}

func (u *fileUseCase) GetFileByID(ctx context.Context, id string) (*domain.UserFile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.fileRepo.GetFileByID(ctx, id)
}

func (f *fileUseCase) UploadFile(ctx context.Context, userID uint, filename, contentType string, data []byte) error {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	// Check if user exists in MySQL
	_, err := f.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Save file in Mongo
	userFile := &domain.UserFile{
		UserID:      userID,
		Filename:    filename,
		ContentType: contentType,
		Data:        data,
	}

	return f.fileRepo.SaveUserFile(ctx, userFile)
}

func (f *fileUseCase) GetFilesByUserID(ctx context.Context, userID uint) ([]*domain.UserFileMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	files, err := f.fileRepo.GetFilesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var meta []*domain.UserFileMeta
	for _, file := range files {
		meta = append(meta, &domain.UserFileMeta{
			ID:       file.ID,
			Filename: file.Filename,
		})
	}

	return meta, nil
}

func (f *fileUseCase) DeleteFilesByUserID(ctx context.Context, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	return f.fileRepo.DeleteFilesByUserID(ctx, userID)
}
