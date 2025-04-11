package controller

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/usecase"
	"boilerplate_go/internal/utils"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context) error
	// Logout(ctx echo.Context) error
	Register(ctx echo.Context) error
	// RefreshToken(ctx echo.Context) error
	Me(ctx echo.Context) error
}

type authController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &authController{authUseCase: authUseCase}
}

// Register implements AuthController.
// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterDTO true "User registration data"
// @Success 201 {object} dto.BaseResponse
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /auth/register [post]
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
// @Summary Login user
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserLoginDTO true "User login data"
// @Success 200 {object} dto.BaseResponse
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /auth/login [post]
func (a *authController) Login(ctx echo.Context) error {
	var userDto dto.UserLoginDTO
	if err := ctx.Bind(&userDto); err != nil {
		return ctx.JSON(400, dto.NewBaseResponse(400, "Invalid request", nil))
	}

	response, _ := a.authUseCase.Login(ctx.Request().Context(), userDto.Username, userDto.Password)

	return ctx.JSON(response.Code, response)
}

// Me implements AuthController.
// @Summary Get user information
// @Description Retrieve current user information using the JWT token
// @Tags Auth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} dto.BaseResponse
// @Failure 401 {object} dto.BaseResponse
// @Failure 404 {object} dto.BaseResponse
// @Router /auth/me [get]
func (a *authController) Me(ctx echo.Context) error {
	username, _ := utils.ExtractUsernameFromToken(ctx.Request().Header.Get("Authorization"))
	response, _ := a.authUseCase.Me(ctx.Request().Context(), username)

	return ctx.JSON(response.Code, response)
}
