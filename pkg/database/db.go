package database

import (
	"fmt"
	"log"
	"os"

	"github.com/rjndrkha/gsitest/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.UserPocket{}, &domain.Income{}, &domain.Expense{})
	if err != nil {
		log.Fatal("Gagal migrasi:", err)
	}

	fmt.Println("Database terkoneksi & Migrasi berhasil!")

	SeedUsers(db)

	return db
}
