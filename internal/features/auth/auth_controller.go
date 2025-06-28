package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/backend-brighted/pkg/utils"
)

type AuthController struct {
	Service AuthService
}

func NewAuthController(s AuthService) *AuthController {
	return &AuthController{Service: s}
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	user, err := controller.Service.Register(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(UserResponse{
		ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role,
	})
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := ac.Service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
	})
}

func (controller *AuthController) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("user").(*utils.JWTClaims)
	user, err := controller.Service.GetProfile(claims.ID)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

func (controller *AuthController) UpdateProfile(c *fiber.Ctx) error {
	claims := c.Locals("user").(*utils.JWTClaims)
	type Req struct {
		Name     string  `json:"name"`
		PhotoURL *string `json:"photo_url"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := controller.Service.UpdateProfile(claims.ID, req.Name, req.PhotoURL); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"message": "Profile updated"})
}
