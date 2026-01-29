package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rjndrkha/gsitest/internal/usecase"
)

type TransactionHandler struct {
	Usecase usecase.TransactionUsecase
}

func (h *TransactionHandler) CreateIncome(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req struct {
		PocketID string `json:"pocket_id"`
		Amount   int64  `json:"amount"`
		Notes    string `json:"notes"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "error": true, "message": "Invalid request"})
	}

	income, balance, err := h.Usecase.CreateIncome(userID, req.PocketID, req.Amount, req.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": 500, "error": true, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": 200, "error": false, "message": "Berhasil menambahkan income.",
		"data": fiber.Map{
			"id": income.ID, "pocket_id": income.PocketID, "current_balance": balance,
		},
	})
}

func (h *TransactionHandler) CreateExpense(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req struct {
		PocketID string `json:"pocket_id"`
		Amount   int64  `json:"amount"`
		Notes    string `json:"notes"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "error": true, "message": "Invalid request"})
	}

	expense, balance, err := h.Usecase.CreateExpense(userID, req.PocketID, req.Amount, req.Notes)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "error": true, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": 200, "error": false, "message": "Berhasil menambahkan expense.",
		"data": fiber.Map{
			"id": expense.ID, "pocket_id": expense.PocketID, "current_balance": balance,
		},
	})
}
