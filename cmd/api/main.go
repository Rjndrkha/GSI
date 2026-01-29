package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rjndrkha/gsitest/internal/handler"
	"github.com/rjndrkha/gsitest/internal/usecase"
	"github.com/rjndrkha/gsitest/pkg/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitDB()
	authUsecase := usecase.AuthUsecase{DB: db}
	authHandler := handler.AuthHandler{Usecase: authUsecase}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API GSI Running!")
	})
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	app.Listen(":8000")
}
