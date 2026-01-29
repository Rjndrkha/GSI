package database

import (
	"github.com/google/uuid"
	"github.com/rjndrkha/gsitest/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	users := []domain.User{
		{
			ID:       uuid.New(),
			FullName: "Rajendra",
			Email:    "rajendra@globalservice.co.id",
			Password: string(hashedPassword),
		},
		{
			ID:       uuid.New(),
			FullName: "Rakha",
			Email:    "rakha@globalservice.co.id",
			Password: string(hashedPassword),
		},
	}

	for _, u := range users {
		db.FirstOrCreate(&u, domain.User{Email: u.Email})
	}
}
