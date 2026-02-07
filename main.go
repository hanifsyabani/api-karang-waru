package main

import (
	"api-karang-waru/config"
	"api-karang-waru/handlers"
	"api-karang-waru/helpers"
	"api-karang-waru/middlewares"
	_ "api-karang-waru/docs" 
	"api-karang-waru/repositories"
	"api-karang-waru/services"
	"log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	subLayananRepository := repositories.NewSubLayananRepository(config.DB)
	riwayatPengajuanRepository := repositories.NewRiwayatPengajuanRepository(config.DB)
	pengajuanLayananRepository := repositories.NewPengajuanLayananRepository(config.DB)

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

	subLayananService := services.NewSubLayananService(subLayananRepository, layananRepository)
	pengajuanLayananService := services.NewPengajuanLayananService(pengajuanLayananRepository, layananRepository, subLayananRepository, riwayatPengajuanRepository)
	riwayatPengajuanService := services.NewRiwayatPengajuanService(riwayatPengajuanRepository)


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

	subLayananHandler := handlers.NewSubLayananHandler(subLayananService)
	pengajuanLayananHandler := handlers.NewPengajuanLayananHandler(pengajuanLayananService)
	riwayatPengajuanHandler := handlers.NewRiwayatPengajuanHandler(riwayatPengajuanService)

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
		base_router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/karang-waru/docs/doc.json")))
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

		auth.GET("/village-profile", profilHandler.GetProfil)
		auth.POST("/village-profile", profilHandler.CreateProfil)
		auth.PUT("/village-profile", profilHandler.UpdateProfil)
		auth.DELETE("/village-profile", profilHandler.DeleteProfil)

		auth.GET("/demographics", demografisHandler.GetDemografis)
		auth.POST("/demographics", demografisHandler.CreateDemografis)
		auth.PUT("/demographics", demografisHandler.UpdateDemografis)
		auth.DELETE("/demographics", demografisHandler.DeleteDemografis)

		auth.GET("/history-village", sejarahHandler.GetSejarah)
		auth.POST("/history-village", sejarahHandler.CreateSejarah)
		auth.PUT("/history-village", sejarahHandler.UpdateSejarah)
		auth.DELETE("/history-village", sejarahHandler.DeleteSejarah)

		auth.GET("/visi-misi", visiMisiHandler.GetVisiMisi)
		auth.POST("/visi-misi", visiMisiHandler.CreateVisiMisi)
		auth.PUT("/visi-misi", visiMisiHandler.UpdateVisiMisi)
		auth.DELETE("/visi-misi", visiMisiHandler.DeleteVisiMisi)

		auth.GET("/news", beritaHandler.GetBerita)
		auth.GET("/news/:id", beritaHandler.GetBeritaByID)
		auth.GET("/news/slug/:slug", beritaHandler.GetBeritaBySlug)
		auth.POST("/news", beritaHandler.CreateBerita)
		auth.PUT("/news/:id", beritaHandler.UpdateBerita)
		auth.DELETE("/news/:id", beritaHandler.DeleteBerita)
		auth.GET("/news/category/count", beritaHandler.GetNewsByCategory)

		auth.GET("/umkm", umkmHandler.GetAllUmkm)
		auth.GET("/umkm/:id", umkmHandler.GetUmkmByID)
		auth.GET("/umkm/slug/:slug", umkmHandler.GetUmkmBySlug)
		auth.GET("/umkm/count-status", umkmHandler.GetCountStatus)
		auth.POST("/umkm", umkmHandler.CreateUmkm)
		auth.PUT("/umkm/:id", umkmHandler.UpdateUmkm)
		auth.DELETE("/umkm/:id", umkmHandler.DeleteUmkm)

		auth.GET("/services", layananHandler.GetAllLayanan)
		auth.GET("/service/:id", layananHandler.GetLayananByID)
		auth.GET("/service/slug/:slug", layananHandler.GetLayananBySlug)
		auth.POST("/service", layananHandler.CreateLayanan)
		auth.PUT("/service/:id", layananHandler.UpdateLayanan)
		auth.DELETE("/service/:id", layananHandler.DeleteLayanan)

		auth.GET("/sub-service", subLayananHandler.GetAllSubLayanan)
		auth.GET("/sub-service/:id", subLayananHandler.GetSubLayananByID)
		auth.POST("/sub-service", subLayananHandler.CreateSubLayanan)
		auth.PUT("/sub-service/:id", subLayananHandler.UpdateSubLayanan)
		auth.DELETE("/sub-service/:id", subLayananHandler.DeleteSubLayanan)
		auth.GET("/sub-service/submissions", pengajuanLayananHandler.GetAllPengajuan)
		auth.POST("/sub-service/submission", pengajuanLayananHandler.CreatePengajuan)
		auth.GET("/sub-service/submission/:id", pengajuanLayananHandler.GetPengajuanByID) 
		auth.GET("/sub-service/submission/nik/:nik", pengajuanLayananHandler.GetPengajuanByNIK) 
		auth.GET("/sub-service/submission/surat/:nomor_surat", pengajuanLayananHandler.GetPengajuanByNomorSurat)
		auth.PUT("/sub-service/submission/:id", pengajuanLayananHandler.UpdatePengajuan)
		auth.PUT("/sub-service/submission/status/:id", pengajuanLayananHandler.UpdateStatusPengajuan)
		auth.PUT("/sub-service/submission/approve/:id", pengajuanLayananHandler.ApprovePengajuan)
		auth.PUT("/sub-service/submission/reject/:id", pengajuanLayananHandler.RejectPengajuan)
		auth.DELETE("/sub-service/submission/:id", pengajuanLayananHandler.DeletePengajuan)
		auth.GET("/sub-service/statistics", pengajuanLayananHandler.GetStatisticsByStatus)
		auth.GET("/sub-service/history/all/:pengajuan_id", riwayatPengajuanHandler.GetAllRiwayat)
		auth.GET("/sub-service/history/:pengajuan_id", riwayatPengajuanHandler.GetAllRiwayat)


		auth.GET("/apbd", apbdHandler.GetApbd)
		auth.GET("/apbd/:id", apbdHandler.GetApbdByID)
		auth.POST("/apbd", apbdHandler.CreateApbd)
		auth.PUT("/apbd/:id", apbdHandler.UpdateApbd)
		auth.DELETE("/apbd/:id", apbdHandler.DeleteApbd)

		auth.GET("/residents", pendudukHandler.GetAllPenduduk)
		auth.GET("/resident/:id", pendudukHandler.GetPendudukByID)
		auth.POST("/resident", pendudukHandler.CreatePenduduk)
		auth.GET("/resident/count", pendudukHandler.CountPenduduk)
		auth.PUT("/resident/:id", pendudukHandler.UpdatePenduduk)
		auth.DELETE("/resident/:id", pendudukHandler.DeletePenduduk)
		// Pendidikan - Program
		auth.GET("/education/program", pendidikanHandler.GetAllProgram)
		auth.GET("/education/program/:id", pendidikanHandler.GetProgramByID)
		auth.POST("/education/program", pendidikanHandler.CreateProgram)
		auth.PUT("/education/program/:id", pendidikanHandler.UpdateProgram)
		auth.DELETE("/education/program/:id", pendidikanHandler.DeleteProgram)

		// Pendidikan - Lembaga
		auth.GET("/education/institution", pendidikanHandler.GetAllLembaga)
		auth.GET("/education/institution/:id", pendidikanHandler.GetLembagaByID)
		auth.POST("/education/institution", pendidikanHandler.CreateLembaga)
		auth.PUT("/education/institution/:id", pendidikanHandler.UpdateLembaga)
		auth.DELETE("/education/institution/:id", pendidikanHandler.DeleteLembaga)

		// Pendidikan - Statistik
		auth.GET("/education/statistic", pendidikanHandler.GetAllStatistik)
		auth.GET("/education/statistic/:id", pendidikanHandler.GetStatistikByID)
		auth.POST("/education/statistic", pendidikanHandler.CreateStatistik)
		auth.PUT("/education/statistic/:id", pendidikanHandler.UpdateStatistik)
		auth.DELETE("/education/statistic/:id", pendidikanHandler.DeleteStatistik)

		// Pendidikan - Capaian
		auth.GET("/education/achievements", pendidikanHandler.GetAllCapaian)
		auth.GET("/education/achievements/:id", pendidikanHandler.GetCapaianByID)
		auth.POST("/education/achievements", pendidikanHandler.CreateCapaian)
		auth.PUT("/education/achievements/:id", pendidikanHandler.UpdateCapaian)
		auth.DELETE("/education/achievements/:id", pendidikanHandler.DeleteCapaian)

		// Pendidikan - Dokumentasi
		auth.GET("/education/documentation", pendidikanHandler.GetAllDokumentasi)
		auth.GET("/education/documentation/:id", pendidikanHandler.GetDokumentasiByID)
		auth.POST("/education/documentation", pendidikanHandler.CreateDokumentasi)
		auth.PUT("/education/documentation/:id", pendidikanHandler.UpdateDokumentasi)
		auth.DELETE("/education/documentation/:id", pendidikanHandler.DeleteDokumentasi)

		// Kesehatan - Layanan
		auth.GET("/health/service", kesehatanHandler.GetLayanan)
		auth.GET("/health/service/:id", kesehatanHandler.GetLayananByID)
		auth.POST("/health/service", kesehatanHandler.CreateLayanan)
		auth.PUT("/health/service/:id", kesehatanHandler.UpdateLayanan)
		auth.DELETE("/health/service/:id", kesehatanHandler.DeleteLayanan)
		auth.GET("/health/facility", kesehatanHandler.GetFasilitasKesehatan)
		auth.GET("/health/facility/:id", kesehatanHandler.GetFasilitasKesehatanByID)
		auth.POST("/health/facility", kesehatanHandler.CreateFasilitas)
		auth.PUT("/health/facility/:id", kesehatanHandler.UpdateFasilitasKesehatan)
		auth.DELETE("/health/facility/:id", kesehatanHandler.DeleteFasilitasKesehatan)

	}

	appPort := config.GetEnv("APP_PORT", "8080")
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
