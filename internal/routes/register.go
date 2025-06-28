package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/backend-brighted/config"
	"github.com/ihsankarim/backend-brighted/internal/features/auth"
)

func Register(app *fiber.App) {
	api := app.Group("/api/v1")

	authRepo := auth.NewAuthRepository(config.DB)
	authService := auth.NewAuthService(authRepo)
	authController := auth.NewAuthController(authService)
	RegisterAuthRoutes(api, authController)
}
