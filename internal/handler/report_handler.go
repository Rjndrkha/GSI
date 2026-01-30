package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rjndrkha/gsitest/internal/usecase"
)

type ReportHandler struct {
	Usecase usecase.ReportUsecase
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {
	pocketID := c.Params("id")
	var req struct {
		Type string `json:"type"`
		Date string `json:"date"`
	}
	c.BodyParser(&req)

	fileID := fmt.Sprintf("%s-%d", uuid.New().String(), time.Now().Unix())

	go h.Usecase.GenerateExcelReport(pocketID, req.Type, req.Date, fileID)

	return c.JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"message": "Report sedang dibuat. Silahkan check berkala pada link berikut.",
		"data": fiber.Map{
			"link": fmt.Sprintf("http://localhost:8000/reports/%s", fileID),
		},
	})
}

func (h *ReportHandler) StreamFile(c *fiber.Ctx) error {
	id := c.Params("id")
	filePath := fmt.Sprintf("./reports/%s.xlsx", id)

	return c.Download(filePath)
}
