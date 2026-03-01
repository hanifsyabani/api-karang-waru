package app

import (
	"api-karang-waru/config"
	"api-karang-waru/internal/delivery/http/router"
	"api-karang-waru/pkg/utils"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	config.InitDB()

	r := gin.Default()

	// CORS
	corsConfig := cors.Config{
		AllowOrigins:     utils.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     utils.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     utils.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: utils.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		ExposeHeaders:    utils.ParseEnvList("CORS_EXPOSE_HEADERS"),
		MaxAge:           12 * 60 * 60,
	}
	if len(corsConfig.AllowOrigins) == 0 {
		corsConfig.AllowAllOrigins = true
	}
	r.Use(cors.New(corsConfig))

	// Routes
	router.SetupRoutes(r, config.DB)

	// Start
	appPort := config.GetEnv("APP_PORT", "8080")
	if err := r.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
