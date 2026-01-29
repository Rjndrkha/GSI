package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rjndrkha/gsitest/internal/usecase"
)

type PocketHandler struct {
	Usecase usecase.PocketUsecase
}

func (h *PocketHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req struct {
		Name           string `json:"name"`
		InitialBalance int64  `json:"initial_balance"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "error": true, "message": "Invalid request"})
	}

	id, err := h.Usecase.CreatePocket(userID, req.Name, req.InitialBalance)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": 500, "error": true, "message": "Gagal membuat pocket"})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"message": "Berhasil membuat pocket baru.",
		"data":    fiber.Map{"id": id},
	})
}

func (h *PocketHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	pockets, err := h.Usecase.GetUserPockets(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": 500, "error": true, "message": "Gagal mengambil data"})
	}

	var data []fiber.Map
	for _, p := range pockets {
		data = append(data, fiber.Map{
			"id":              p.ID,
			"name":            p.Name,
			"current_balance": p.Balance,
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"message": "Berhasil.",
		"data":    data,
	})
}
