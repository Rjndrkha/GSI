package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rjndrkha/gsitest/pkg/database"
)

func main() {
	database.InitDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API GSI Running!")
	})

	app.Listen(":8000")
}
