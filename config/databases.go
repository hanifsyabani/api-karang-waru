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
	dsn := GetEnv("SUPABASE_DSN", "")
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DSN tidak diset")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	fmt.Println("✅ Database terkoneksi")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}
	DB = db
}
