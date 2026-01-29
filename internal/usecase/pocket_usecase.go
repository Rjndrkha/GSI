package usecase

import (
	"github.com/google/uuid"
	"github.com/rjndrkha/gsitest/internal/domain"

	"gorm.io/gorm"
)

type PocketUsecase struct {
	DB *gorm.DB
}

func (u *PocketUsecase) CreatePocket(userID string, name string, initialBalance int64) (uuid.UUID, error) {
	uID, _ := uuid.Parse(userID)
	pocket := domain.UserPocket{
		ID:      uuid.New(),
		UserID:  uID,
		Name:    name,
		Balance: initialBalance,
	}

	if err := u.DB.Create(&pocket).Error; err != nil {
		return uuid.Nil, err
	}
	return pocket.ID, nil
}

func (u *PocketUsecase) GetUserPockets(userID string) ([]domain.UserPocket, error) {
	var pockets []domain.UserPocket
	if err := u.DB.Where("user_id = ?", userID).Find(&pockets).Error; err != nil {
		return nil, err
	}
	return pockets, nil
}
