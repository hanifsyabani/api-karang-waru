package main

import (
	"api-karang-waru/internal/app"
	"log"

	"github.com/joho/godotenv"
)

// @title Karang Waru API
// @version 1.0
// @description API Desa Karang Waru
// @host eventual-alika-karang-waru-aabb0b75.koyeb.app
// @BasePath /api/karang-waru
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No env file found, relying on environment variables.")
	}

	app.Run()
}
