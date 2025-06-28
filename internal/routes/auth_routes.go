package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/backend-brighted/internal/features/auth"
	"github.com/ihsankarim/backend-brighted/pkg/middleware"
)

func RegisterAuthRoutes(routes fiber.Router, controller *auth.AuthController) {
	authGroup := routes.Group("/auth")
	authGroup.Post("/register", controller.Register)
	authGroup.Post("/login", controller.Login)

	protected := authGroup.Use(middleware.JWTMiddleware())
	protected.Get("/me", controller.GetProfile)
	protected.Put("/me", controller.UpdateProfile)
}
