package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FullName string    `gorm:"not null" json:"full_name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
}

type UserPocket struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name    string    `gorm:"not null" json:"name"`
	Balance int64     `gorm:"default:0" json:"balance"`
}

type Income struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	PocketID uuid.UUID `gorm:"type:uuid" json:"pocket_id"`
	Amount   int64     `gorm:"not null" json:"amount"`
	Notes    string    `json:"notes"`
}

type Expense struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	PocketID uuid.UUID `gorm:"type:uuid" json:"pocket_id"`
	Amount   int64     `gorm:"not null" json:"amount"`
	Notes    string    `json:"notes"`
}
