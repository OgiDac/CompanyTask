package repository

import (
	"context"
	"errors"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id uint) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	var existing domain.User
	if err := u.db.WithContext(ctx).First(&existing, user.ID).Error; err != nil {
		return err
	}

	// Update fields
	existing.Name = user.Name
	existing.Email = user.Email

	// Save
	if err := u.db.WithContext(ctx).Save(&existing).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id uint) error {
	result := u.db.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
