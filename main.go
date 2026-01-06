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
		log.Println("No env file found, relying on environment variables.")
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
	umkmRepository := repositories.NewUmkmRepository(config.DB)
	layananRepository := repositories.NewLayananRepository(config.DB)
	apbdRepository := repositories.NewApbdRepository(config.DB)
	pendudukRepository := repositories.NewPendudukRepository(config.DB)
	pendidikanRepository := repositories.NewPendidikanRepository(config.DB)
	kesehatanRepository := repositories.NewKesehatanRepository(config.DB)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()
	profilService := services.NewProfilService(profilRepository)
	demografisService := services.NewDemografisService(demografisRepository)
	sejarahService := services.NewSejarahService(sejarahRepository)
	visiMisiService := services.NewVisiMisiService(visiMisiRepository)
	beritaService := services.NewBeritaService(beritaRepository)
	umkmService := services.NewUmkmService(umkmRepository)
	layananService := services.NewLayananService(layananRepository)
	apbdService := services.NewApbdService(apbdRepository)
	pendudukService := services.NewPendudukService(pendudukRepository)
	pendidikanService := services.NewPendidikanService(pendidikanRepository)
	kesehatanService := services.NewKesehatanService(kesehatanRepository)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	profilHandler := handlers.NewProfilDesaHandler(profilService)
	demografisHandler := handlers.NewDemografisHandler(demografisService)
	sejarahHandler := handlers.NewSejarahHandler(sejarahService)
	visiMisiHandler := handlers.NewVisiMisiHandler(visiMisiService)
	beritaHandler := handlers.NewBeritaHandler(beritaService)
	umkmHandler := handlers.NewUmkmHandler(umkmService)
	layananHandler := handlers.NewLayananHandler(layananService)
	apbdHandler := handlers.NewApbdHandler(apbdService)
	pendudukHandler := handlers.NewPendudukHandler(pendudukService)
	pendidikanHandler := handlers.NewPendidikanHandler(pendidikanService)
	kesehatanHandler := handlers.NewKesehatanHandler(kesehatanService)

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
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/me", userHandler.GetProfile)
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
		auth.GET("/berita/slug/:slug", beritaHandler.GetBeritaBySlug)
		auth.POST("/berita", beritaHandler.CreateBerita)
		auth.PUT("/berita/:id", beritaHandler.UpdateBerita)
		auth.DELETE("/berita/:id", beritaHandler.DeleteBerita)
		auth.GET("/umkm", umkmHandler.GetAllUmkm)
		auth.GET("/umkm/:id", umkmHandler.GetUmkmByID)
		auth.GET("/umkm/slug/:slug", umkmHandler.GetUmkmBySlug)
		auth.POST("/umkm", umkmHandler.CreateUmkm)
		auth.PUT("/umkm/:id", umkmHandler.UpdateUmkm)
		auth.DELETE("/umkm/:id", umkmHandler.DeleteUmkm)
		auth.GET("/layanan", layananHandler.GetAllLayanan)
		auth.GET("/layanan/:id", layananHandler.GetLayananByID)
		auth.GET("/layanan/slug/:slug", layananHandler.GetLayananBySlug)
		auth.POST("/layanan", layananHandler.CreateLayanan)
		auth.PUT("/layanan/:id", layananHandler.UpdateLayanan)
		auth.DELETE("/layanan/:id", layananHandler.DeleteLayanan)
		auth.GET("/apbd", apbdHandler.GetApbd)
		auth.GET("/apbd/:id", apbdHandler.GetApbdByID)
		auth.POST("/apbd", apbdHandler.CreateApbd)
		auth.PUT("/apbd/:id", apbdHandler.UpdateApbd)
		auth.DELETE("/apbd/:id", apbdHandler.DeleteApbd)
		auth.GET("/penduduk", pendudukHandler.GetAllPenduduk)
		auth.GET("/penduduk/:id", pendudukHandler.GetPendudukByID)
		auth.POST("/penduduk", pendudukHandler.CreatePenduduk)
		auth.PUT("/penduduk/:id", pendudukHandler.UpdatePenduduk)
		auth.DELETE("/penduduk/:id", pendudukHandler.DeletePenduduk)
		// Pendidikan - Program
		auth.GET("/pendidikan/program", pendidikanHandler.GetAllProgram)
		auth.GET("/pendidikan/program/:id", pendidikanHandler.GetProgramByID)
		auth.POST("/pendidikan/program", pendidikanHandler.CreateProgram)
		auth.PUT("/pendidikan/program/:id", pendidikanHandler.UpdateProgram)
		auth.DELETE("/pendidikan/program/:id", pendidikanHandler.DeleteProgram)

		// Pendidikan - Lembaga
		auth.GET("/pendidikan/lembaga", pendidikanHandler.GetAllLembaga)
		auth.GET("/pendidikan/lembaga/:id", pendidikanHandler.GetLembagaByID)
		auth.POST("/pendidikan/lembaga", pendidikanHandler.CreateLembaga)
		auth.PUT("/pendidikan/lembaga/:id", pendidikanHandler.UpdateLembaga)
		auth.DELETE("/pendidikan/lembaga/:id", pendidikanHandler.DeleteLembaga)

		// Pendidikan - Statistik
		auth.GET("/pendidikan/statistik", pendidikanHandler.GetAllStatistik)
		auth.GET("/pendidikan/statistik/:id", pendidikanHandler.GetStatistikByID)
		auth.POST("/pendidikan/statistik", pendidikanHandler.CreateStatistik)
		auth.PUT("/pendidikan/statistik/:id", pendidikanHandler.UpdateStatistik)
		auth.DELETE("/pendidikan/statistik/:id", pendidikanHandler.DeleteStatistik)

		// Pendidikan - Capaian
		auth.GET("/pendidikan/capaian", pendidikanHandler.GetAllCapaian)
		auth.GET("/pendidikan/capaian/:id", pendidikanHandler.GetCapaianByID)
		auth.POST("/pendidikan/capaian", pendidikanHandler.CreateCapaian)
		auth.PUT("/pendidikan/capaian/:id", pendidikanHandler.UpdateCapaian)
		auth.DELETE("/pendidikan/capaian/:id", pendidikanHandler.DeleteCapaian)

		// Pendidikan - Dokumentasi
		auth.GET("/pendidikan/dokumentasi", pendidikanHandler.GetAllDokumentasi)
		auth.GET("/pendidikan/dokumentasi/:id", pendidikanHandler.GetDokumentasiByID)
		auth.POST("/pendidikan/dokumentasi", pendidikanHandler.CreateDokumentasi)
		auth.PUT("/pendidikan/dokumentasi/:id", pendidikanHandler.UpdateDokumentasi)
		auth.DELETE("/pendidikan/dokumentasi/:id", pendidikanHandler.DeleteDokumentasi)

		// Kesehatan - Layanan
		auth.GET("/kesehatan/layanan", kesehatanHandler.GetLayanan)
		auth.GET("/kesehatan/layanan/:id", kesehatanHandler.GetLayananByID)
		auth.POST("/kesehatan/layanan", kesehatanHandler.CreateLayanan)
		auth.PUT("/kesehatan/layanan/:id", kesehatanHandler.UpdateLayanan)
		auth.DELETE("/kesehatan/layanan/:id", kesehatanHandler.DeleteLayanan)
		auth.GET("/kesehatan/fasilitas", kesehatanHandler.GetFasilitasKesehatan)
		auth.GET("/kesehatan/fasilitas/:id", kesehatanHandler.GetFasilitasKesehatanByID)
		auth.POST("/kesehatan/fasilitas", kesehatanHandler.CreateFasilitas)
		auth.PUT("/kesehatan/fasilitas/:id", kesehatanHandler.UpdateFasilitasKesehatan)
		auth.DELETE("/kesehatan/fasilitas/:id", kesehatanHandler.DeleteFasilitasKesehatan)

	}

	appPort := config.GetEnv("APP_PORT", "8080")
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
