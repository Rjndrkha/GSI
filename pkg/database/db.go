package database

import (
	"fmt"
	"log"

	"github.com/rjndrkha/gsitest/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=golearn port=5432 sslmode=disable"
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
