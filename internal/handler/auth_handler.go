package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rjndrkha/gsitest/internal/usecase"
)

type AuthHandler struct {
	Usecase usecase.AuthUsecase
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "error": true, "message": "Invalid request"})
	}

	token, err := h.Usecase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": 401, "error": true, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"message": "Berhasil login.",
		"data": fiber.Map{
			"token": token,
		},
	})
}

func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.Usecase.GetProfile(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": 404, "error": true, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"message": "Berhasil login.",
		"data": fiber.Map{
			"full_name": user.FullName,
			"email":     user.Email,
		},
	})
}
