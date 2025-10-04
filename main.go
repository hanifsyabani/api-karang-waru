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
	profilRepository := repositories.NewProfilRepository(config.DB)
	demografisRepository := repositories.NewDemografisRepository(config.DB)
	sejarahRepository := repositories.NewSejarahRepository(config.DB)
	visiMisiRepository := repositories.NewVisiMisiRepository(config.DB)
	beritaRepository := repositories.NewBeritaRepository(config.DB)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()
	profilService := services.NewProfilService(profilRepository)
	demografisService := services.NewDemografisService(demografisRepository)
	sejarahService := services.NewSejarahService(sejarahRepository)
	visiMisiService := services.NewVisiMisiService(visiMisiRepository)
	beritaService := services.NewBeritaService(beritaRepository)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	profilHandler := handlers.NewProfilDesaHandler(profilService)
	demografisHandler := handlers.NewDemografisHandler(demografisService)
	sejarahHandler := handlers.NewSejarahHandler(sejarahService)
	visiMisiHandler := handlers.NewVisiMisiHandler(visiMisiService)
	beritaHandler := handlers.NewBeritaHandler(beritaService)

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

	api := router.Group("/api")
	base_router := api.Group("/karang-waru")
	{
		base_router.POST("/register", authHandler.Register)
		base_router.POST("/login", authHandler.Login)
		base_router.POST("/logout", authHandler.Logout)
	}

	// protected
	auth := api.Group("/karang-waru")
	auth.Use(middlewares.JWTAuth())
	{
		auth.GET("/profile", userHandler.GetProfile)
		auth.GET("/users", userHandler.GetUsers)
		auth.GET("/users/:id", userHandler.GetUser)
		auth.POST("/users", userHandler.CreateUser)
		auth.PUT("/users/:id", userHandler.UpdateUser)
		auth.DELETE("/users/:id", userHandler.DeleteUser)
		auth.GET("/profil-desa", profilHandler.GetProfil)
		auth.POST("/profil-desa", profilHandler.CreateProfil)
		auth.PUT("/profil-desa", profilHandler.UpdateProfil)
		auth.DELETE("/profil-desa", profilHandler.DeleteProfil)
		auth.GET("/demografis", demografisHandler.GetDemografis)
		auth.POST("/demografis", demografisHandler.CreateDemografis)
		auth.PUT("/demografis", demografisHandler.UpdateDemografis)
		auth.DELETE("/demografis", demografisHandler.DeleteDemografis)
		auth.GET("/sejarah", sejarahHandler.GetSejarah)
		auth.POST("/sejarah", sejarahHandler.CreateSejarah)
		auth.PUT("/sejarah", sejarahHandler.UpdateSejarah)
		auth.DELETE("/sejarah", sejarahHandler.DeleteSejarah)
		auth.GET("/visi-misi", visiMisiHandler.GetVisiMisi)
		auth.POST("/visi-misi", visiMisiHandler.CreateVisiMisi)
		auth.PUT("/visi-misi", visiMisiHandler.UpdateVisiMisi)
		auth.DELETE("/visi-misi", visiMisiHandler.DeleteVisiMisi)
		auth.GET("/berita", beritaHandler.GetBerita)
		auth.GET("/berita/:id", beritaHandler.GetBeritaByID)
		auth.POST("/berita", beritaHandler.CreateBerita)
		auth.PUT("/berita/:id", beritaHandler.UpdateBerita)
		auth.DELETE("/berita/:id", beritaHandler.DeleteBerita)
	}

	appPort := config.GetEnv("APP_PORT", "8080")
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
