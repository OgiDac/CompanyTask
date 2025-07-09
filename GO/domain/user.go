package domain

import "context"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:255" json:"name"`
	Email    string `gorm:"size:255;unique" json:"email"`
	Password string `gorm:"password" json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UpdateRequest struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserUseCase interface {
	GetAllUsers(c context.Context) ([]*UserResponse, error)
	CreateUser(c context.Context, user SignUpRequest) (accessToken string, refreshToken string, err error)
	UpdateUser(c context.Context, user UpdateRequest) error
	Login(ctx context.Context, request LoginRequest) (accessToken string, refreshToken string, err error)
	DeleteUser(ctx context.Context, id uint) error
}
