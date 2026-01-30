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

	pocketUsecase := usecase.PocketUsecase{DB: db}
	pocketHandler := handler.PocketHandler{Usecase: pocketUsecase}

	transactionUsecase := usecase.TransactionUsecase{DB: db}
	transactionHandler := handler.TransactionHandler{Usecase: transactionUsecase}

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Get("/profile", middleware.JWTMiddleware, authHandler.GetProfile)

	pockets := api.Group("/pockets", middleware.JWTMiddleware)
	pockets.Post("/", pocketHandler.Create)
	pockets.Get("/", pocketHandler.List)
	pockets.Get("/total-balance", pocketHandler.GetTotal)

	api.Post("/incomes", middleware.JWTMiddleware, transactionHandler.CreateIncome)
	api.Post("/expenses", middleware.JWTMiddleware, transactionHandler.CreateExpense)
}
