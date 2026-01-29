package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rjndrkha/gsitest/internal/domain"

	"gorm.io/gorm"
)

type TransactionUsecase struct {
	DB *gorm.DB
}

func (u *TransactionUsecase) CreateIncome(userID, pocketID string, amount int64, notes string) (domain.Income, int64, error) {
	var income domain.Income
	var currentBalance int64
	uID, _ := uuid.Parse(userID)
	pID, _ := uuid.Parse(pocketID)

	err := u.DB.Transaction(func(tx *gorm.DB) error {

		income = domain.Income{
			ID: uuid.New(), UserID: uID, PocketID: pID, Amount: amount, Notes: notes,
		}
		if err := tx.Create(&income).Error; err != nil {
			return err
		}

		var pocket domain.UserPocket
		if err := tx.Model(&pocket).Where("id = ? AND user_id = ?", pID, uID).
			Update("balance", gorm.Expr("balance + ?", amount)).First(&pocket).Error; err != nil {
			return errors.New("pocket tidak ditemukan")
		}
		currentBalance = pocket.Balance
		return nil
	})

	return income, currentBalance, err
}

func (u *TransactionUsecase) CreateExpense(userID, pocketID string, amount int64, notes string) (domain.Expense, int64, error) {
	var expense domain.Expense
	var currentBalance int64
	uID, _ := uuid.Parse(userID)
	pID, _ := uuid.Parse(pocketID)

	err := u.DB.Transaction(func(tx *gorm.DB) error {
		var pocket domain.UserPocket
		if err := tx.Where("id = ? AND user_id = ?", pID, uID).First(&pocket).Error; err != nil {
			return errors.New("pocket tidak ditemukan")
		}

		if pocket.Balance < amount {
			return errors.New("saldo tidak cukup")
		}

		expense = domain.Expense{
			ID: uuid.New(), UserID: uID, PocketID: pID, Amount: amount, Notes: notes,
		}
		if err := tx.Create(&expense).Error; err != nil {
			return err
		}

		tx.Model(&pocket).Update("balance", pocket.Balance-amount)
		currentBalance = pocket.Balance - amount
		return nil
	})

	return expense, currentBalance, err
}
