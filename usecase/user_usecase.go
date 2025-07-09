package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/OgiDac/CompanyTask/repository"
	"github.com/OgiDac/CompanyTask/utils"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *userUseCase) GetAllUsers(c context.Context) ([]*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var userResponse []*domain.UserResponse
	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userResponse = append(userResponse, &domain.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponse, nil
}

func (u *userUseCase) CreateUser(c context.Context, user domain.SignUpRequest) (string, string, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", "", err
	}

	user.Password = string(encryptedPassword)

	signUpUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err = u.userRepository.CreateUser(ctx, signUpUser)
	if err != nil {
		return "", "", err
	}

	accessToken, err := utils.CreateAccessToken(signUpUser, "secret", 5)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(signUpUser, "secret", 5)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *userUseCase) UpdateUser(c context.Context, req domain.UpdateRequest) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	updatedUser := &domain.User{
		ID:    uint(req.Id),
		Name:  req.Name,
		Email: req.Email,
	}

	return u.userRepository.UpdateUser(ctx, updatedUser)
}

func (u *userUseCase) Login(ctx context.Context, request domain.LoginRequest) (string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return "", "", errors.New("user does not exist")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return "", "", errors.New("invalid password")
	}

	accessToken, err := utils.CreateAccessToken(user, "secret", 5)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(user, "secret", 5)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.DeleteUser(ctx, id)
}
