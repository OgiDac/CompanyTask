package domain

import "context"

type UserFile struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	UserID      uint   `bson:"userId" json:"userId"`
	Filename    string `bson:"filename" json:"filename"`
	ContentType string `bson:"contentType" json:"contentType"`
	Data        []byte `bson:"data" json:"-"`
}

type UserFileMeta struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}

type FileUseCase interface {
	UploadFile(ctx context.Context, userID uint, filename, contentType string, data []byte) error
	GetFileByID(ctx context.Context, id string) (*UserFile, error)
	GetFilesByUserID(ctx context.Context, userID uint) ([]*UserFileMeta, error)
	DeleteFilesByUserID(ctx context.Context, userID uint) error
}
