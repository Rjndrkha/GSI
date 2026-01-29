package usecase

import (
	"errors"

	"github.com/rjndrkha/gsitest/internal/domain"
	"github.com/rjndrkha/gsitest/pkg/util"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB *gorm.DB
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	var user domain.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("email atau password salah")
	}

	if err := util.ComparePassword(user.Password, password); err != nil {
		return "", errors.New("email atau password salah")
	}

	return util.GenerateToken(user.ID.String())
}

func (u *AuthUsecase) GetProfile(userID string) (domain.User, error) {
	var user domain.User
	if err := u.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, errors.New("user tidak ditemukan")
	}
	return user, nil
}
