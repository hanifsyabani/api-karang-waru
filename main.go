package main

import (
	"api-karang-waru/config"
	"api-karang-waru/handlers"
	"api-karang-waru/helpers"
	"api-karang-waru/middlewares"
	"api-karang-waru/repositories"
	"api-karang-waru/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	config.InitDB()

	mainHandler := handlers.NewMainHandler()
	healthHandler := handlers.NewHealthHandler()
	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     helpers.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helpers.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helpers.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helpers.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		ExposeHeaders:    helpers.ParseEnvList("CORS_EXPOSE_HEADERS"),
		MaxAge:           12 * 60 * 60,
	}

	if len(corsConfig.AllowOrigins) == 0 {
		corsConfig.AllowAllOrigins = true
	}

	router.Use(cors.New(corsConfig))
	router.GET("/", mainHandler.MainHandler)
	router.GET("/health", healthHandler.HealthCheck)
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	auth := router.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/users", userHandler.GetUsers)
		auth.GET("/users/:id", userHandler.GetUser)
		auth.POST("/users", userHandler.CreateUser)
		auth.PUT("/users/:id", userHandler.UpdateUser)
		auth.DELETE("/users/:id", userHandler.DeleteUser)
	}

	appPort := config.GetEnv("APP_PORT", "8080")
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
