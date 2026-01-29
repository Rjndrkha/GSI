package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rjndrkha/gsitest/internal/routes"
	"github.com/rjndrkha/gsitest/pkg/database"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitDB()

	app := fiber.New()
	routes.SetupRoutes(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API GSI Running!")
	})

	app.Listen(":8000")
}
