package usecase

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/model"
	"boilerplate_go/internal/repository"
	"boilerplate_go/internal/utils"
	"context"
	"errors"
	"time"
)

type AuthUseCase interface {
	Register(ctx context.Context, userDto *dto.UserRegisterDTO) (dto.BaseResponse, error)
	Login(ctx context.Context, username, password string) (dto.BaseResponse, error)
	Me(ctx context.Context, username string) (dto.BaseResponse, error)
}

type authUseCase struct {
	repo repository.BaseRepository[model.User]
}

func NewAuthUseCase(repo repository.BaseRepository[model.User]) AuthUseCase {
	return &authUseCase{repo: repo}
}

func (a *authUseCase) Register(ctx context.Context, userDto *dto.UserRegisterDTO) (dto.BaseResponse, error) {
	// Check if user already exists
	existingUser, err := a.repo.First("username = ?", userDto.Username)
	if err == nil && existingUser != nil {
		return dto.NewBaseResponse(400, "User already exists", nil), errors.New("user already exists")
	}

	// Hash the provided password.
	hashedPassword, err := utils.HashedPassword(userDto.Password)
	if err != nil {
		return dto.NewBaseResponse(500, "Failed to hash password", nil), err
	}

	// Create a new user entity.
	newUser := model.User{
		ID:           utils.GenerateUserID(userDto.Username),
		Username:     userDto.Username,
		PasswordHash: hashedPassword,
		Email:        userDto.Email,
		CreatedAt:    time.Now(),
	}

	if err := a.repo.Create(&newUser); err != nil {
		return dto.NewBaseResponse(500, "Failed to create user", nil), err
	}

	return dto.NewBaseResponse(201, "Registration successful", nil), nil
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(ctx context.Context, username string, password string) (dto.BaseResponse, error) {
	existingUser, err := a.repo.First("username = ?", username)
	if err != nil || existingUser == nil {
		return dto.NewBaseResponse(400, "Invalid user info.", nil), err
	}

	if !utils.CheckPasswordHash(password, existingUser.PasswordHash) {
		return dto.NewBaseResponse(400, "Invalid user info.", nil), nil
	}

	tokenString, expired, err := utils.JWTSecret(username)
	if err != nil {
		return dto.NewBaseResponse(500, "Failed to generate token", nil), err
	}

	loginResponse := dto.UserLoginResponse{
		Username: username,
		Role:     "todo",
		Token:    tokenString,
		Expiry:   expired,
	}

	return dto.NewBaseResponse(200, "Login successful", loginResponse), nil
}

// Me implements AuthUseCase.
func (a *authUseCase) Me(ctx context.Context, username string) (dto.BaseResponse, error) {
	existingUser, err := a.repo.First("username = ?", username)
	if err != nil || existingUser == nil {
		return dto.NewBaseResponse(400, "Invalid user info.", nil), err
	}

	userDto := dto.UserDTO{
		ID:        existingUser.ID,
		Username:  existingUser.Username,
		Email:     existingUser.Email,
		CreatedAt: existingUser.CreatedAt.Format(time.RFC3339),
	}

	return dto.NewBaseResponse(200, "", userDto), nil
}
