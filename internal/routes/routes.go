package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rjndrkha/gsitest/internal/handler"
	"github.com/rjndrkha/gsitest/internal/usecase"
	"github.com/rjndrkha/gsitest/pkg/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	authUsecase := usecase.AuthUsecase{DB: db}
	authHandler := handler.AuthHandler{Usecase: authUsecase}

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Get("/profile", middleware.JWTMiddleware, authHandler.GetProfile)

}
