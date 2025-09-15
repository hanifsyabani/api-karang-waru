package config

import (
	"api-karang-waru/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := GetEnv("DATABASE_URL", "")
	if dsn == "" {
		log.Fatal("Db url not found")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	fmt.Println("✅ DB Connected")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}
	DB = db
}
