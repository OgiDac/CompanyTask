package repository

import (
	"context"
	"errors"

	"github.com/OgiDac/CompanyTask/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepository interface {
	SaveUserFile(ctx context.Context, file *domain.UserFile) error
	GetFileByID(ctx context.Context, id string) (*domain.UserFile, error)
	GetFilesByUserID(ctx context.Context, userID uint) ([]*domain.UserFile, error)
	DeleteFilesByUserID(ctx context.Context, userID uint) error
}

type fileRepository struct {
	collection *mongo.Collection
}

func NewFileRepository(db *mongo.Database) FileRepository {
	return &fileRepository{
		collection: db.Collection("user_files"),
	}
}

func (f *fileRepository) SaveUserFile(ctx context.Context, file *domain.UserFile) error {
	_, err := f.collection.InsertOne(ctx, file)
	return err
}

func (r *fileRepository) GetFileByID(ctx context.Context, id string) (*domain.UserFile, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	var result domain.UserFile
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *fileRepository) GetFilesByUserID(ctx context.Context, userID uint) ([]*domain.UserFile, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []*domain.UserFile
	for cursor.Next(ctx) {
		var f domain.UserFile
		if err := cursor.Decode(&f); err != nil {
			continue
		}
		files = append(files, &f)
	}

	return files, nil
}

func (r *fileRepository) DeleteFilesByUserID(ctx context.Context, userID uint) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"userId": userID})
	return err
}
