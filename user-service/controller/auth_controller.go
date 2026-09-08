package controller

import (
	"pos/user-service/controller/request"
	"pos/user-service/controller/response"
	"pos/user-service/pkg/conv"
	"pos/user-service/pkg/validator"
	"pos/user-service/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type AuthControllerInterface interface {
	Login(c fiber.Ctx) error
}

type AuthController struct {
	AuthService usecase.UserUsecaseInterface
}

// Login implementasi AuthControllerInterface
func (a *AuthController) Login(c fiber.Ctx) error {
	ctx := c.Context()
	var loginRequest request.LoginRequest
	if err := c.Bind().Body(&loginRequest); err != nil {
		log.Errorf("[AuthController.Login] Login - 1: %v", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(loginRequest); err != nil {
		log.Errorf("[AuthController.Login] Login - 2: %v", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	user, err := a.AuthService.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		log.Errorf("[AuthController.Login] Login - 3: %v", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	if user == nil {
		log.Errorf("[AuthController.Login] Login - 4: user not found for email %s", loginRequest.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	isSame := conv.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isSame {
		log.Errorf("[AuthController.Login] Login - 5: password mismatch for email %s", loginRequest.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// satu user satu role

	roleName := ""
	if len(user.Roles) > 0 {
		roleName = user.Roles[0].Name
	}

	loginResp := response.LoginResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Role: roleName,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"data":    loginResp,
	})
}

func NewAuthController(authService usecase.UserUsecaseInterface) AuthControllerInterface {
	return &AuthController{
		AuthService: authService,
	}
}