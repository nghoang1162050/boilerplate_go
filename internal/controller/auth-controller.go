package controller

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/usecase"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context) error
	// Logout(ctx echo.Context) error
	Register(ctx echo.Context) error
	// RefreshToken(ctx echo.Context) error
	// GetProfile(ctx echo.Context) error
}

type authController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &authController{authUseCase: authUseCase}
}

// Register implements AuthController.
func (a *authController) Register(ctx echo.Context) error {
	var userDto dto.UserRegisterDTO
	if err := ctx.Bind(&userDto); err != nil {
		return ctx.JSON(400, dto.NewBaseResponse(400, "Invalid request", nil))
	}

	// TODO: Uncomment and implement validation logic
	// if err := ctx.Validate(&userDto); err != nil {
	// 	return ctx.JSON(400, dto.NewBaseResponse(400, "Validation failed", nil))
	// }

	response, _ := a.authUseCase.Register(ctx.Request().Context(), &userDto)

	return ctx.JSON(response.Code, response)
}

// Login implements AuthController.
func (a *authController) Login(ctx echo.Context) error {
	var userDto dto.UserLoginDTO
	if err := ctx.Bind(&userDto); err != nil {
		return ctx.JSON(400, dto.NewBaseResponse(400, "Invalid request", nil))
	}

	response,_ := a.authUseCase.Login(ctx.Request().Context(), userDto.Username, userDto.Password)

	return ctx.JSON(response.Code, response)
}
