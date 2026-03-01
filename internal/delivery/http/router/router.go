package router

import (
	"api-karang-waru/internal/delivery/http/handler"
	"api-karang-waru/internal/delivery/http/middleware"
	"api-karang-waru/internal/usecase"

	_ "api-karang-waru/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// ── Repositories ──
	userRepo := usecase.NewUserRepository(db)
	profilRepo := usecase.NewProfilRepository(db)
	demografisRepo := usecase.NewDemografisRepository(db)
	sejarahRepo := usecase.NewSejarahRepository(db)
	visiMisiRepo := usecase.NewVisiMisiRepository(db)
	beritaRepo := usecase.NewBeritaRepository(db)
	umkmRepo := usecase.NewUmkmRepository(db)
	layananRepo := usecase.NewLayananRepository(db)
	apbdRepo := usecase.NewApbdRepository(db)
	pendudukRepo := usecase.NewPendudukRepository(db)
	pendidikanRepo := usecase.NewPendidikanRepository(db)
	kesehatanRepo := usecase.NewKesehatanRepository(db)
	subLayananRepo := usecase.NewSubLayananRepository(db)
	riwayatRepo := usecase.NewRiwayatPengajuanRepository(db)
	pengajuanRepo := usecase.NewPengajuanLayananRepository(db)

	// ── Use Cases (Services) ──
	userUsecase := usecase.NewUserService(userRepo)
	authUsecase := usecase.NewAuthService()
	profilUsecase := usecase.NewProfilService(profilRepo)
	demografisUsecase := usecase.NewDemografisService(demografisRepo)
	sejarahUsecase := usecase.NewSejarahService(sejarahRepo)
	visiMisiUsecase := usecase.NewVisiMisiService(visiMisiRepo)
	beritaUsecase := usecase.NewBeritaService(beritaRepo)
	umkmUsecase := usecase.NewUmkmService(umkmRepo)
	layananUsecase := usecase.NewLayananService(layananRepo)
	apbdUsecase := usecase.NewApbdService(apbdRepo)
	pendudukUsecase := usecase.NewPendudukService(pendudukRepo)
	pendidikanUsecase := usecase.NewPendidikanService(pendidikanRepo)
	kesehatanUsecase := usecase.NewKesehatanService(kesehatanRepo)
	subLayananUsecase := usecase.NewSubLayananService(subLayananRepo, layananRepo)
	pengajuanUsecase := usecase.NewPengajuanLayananService(pengajuanRepo, layananRepo, subLayananRepo, riwayatRepo)
	riwayatUsecase := usecase.NewRiwayatPengajuanService(riwayatRepo)

	// ── Handlers ──
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	profilHandler := handler.NewProfilDesaHandler(profilUsecase)
	demografisHandler := handler.NewDemografisHandler(demografisUsecase)
	sejarahHandler := handler.NewSejarahHandler(sejarahUsecase)
	visiMisiHandler := handler.NewVisiMisiHandler(visiMisiUsecase)
	beritaHandler := handler.NewBeritaHandler(beritaUsecase)
	umkmHandler := handler.NewUmkmHandler(umkmUsecase)
	layananHandler := handler.NewLayananHandler(layananUsecase)
	apbdHandler := handler.NewApbdHandler(apbdUsecase)
	pendudukHandler := handler.NewPendudukHandler(pendudukUsecase)
	pendidikanHandler := handler.NewPendidikanHandler(pendidikanUsecase)
	kesehatanHandler := handler.NewKesehatanHandler(kesehatanUsecase)
	subLayananHandler := handler.NewSubLayananHandler(subLayananUsecase)
	pengajuanHandler := handler.NewPengajuanLayananHandler(pengajuanUsecase)
	riwayatHandler := handler.NewRiwayatPengajuanHandler(riwayatUsecase)
	dashboardHandler := handler.NewDashboardHandler()
	healthHandler := handler.NewHealthHandler()

	// ── Root routes ──
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Desa Karang Waru"})
	})
	r.GET("/health", healthHandler.HealthCheck)

	// ── Public routes ──
	api := r.Group("/api")
	public := api.Group("/karang-waru")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.POST("/logout", authHandler.Logout)
		public.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/karang-waru/docs/doc.json")))
	}

	// ── Protected routes ──
	protected := api.Group("/karang-waru")
	protected.Use(middleware.AuthMiddleware())
	{
		// User
		protected.GET("/me", userHandler.GetProfile)
		protected.GET("/users", userHandler.GetUsers)
		protected.GET("/users/:id", userHandler.GetUser)
		protected.PUT("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)

		// Village Profile
		protected.GET("/village-profile", profilHandler.GetProfil)
		protected.POST("/village-profile", profilHandler.CreateProfil)
		protected.PUT("/village-profile", profilHandler.UpdateProfil)
		protected.DELETE("/village-profile", profilHandler.DeleteProfil)

		// Demographics
		protected.GET("/demographics", demografisHandler.GetDemografis)
		protected.POST("/demographics", demografisHandler.CreateDemografis)
		protected.PUT("/demographics", demografisHandler.UpdateDemografis)
		protected.DELETE("/demographics", demografisHandler.DeleteDemografis)

		// History
		protected.GET("/history-village", sejarahHandler.GetSejarah)
		protected.POST("/history-village", sejarahHandler.CreateSejarah)
		protected.PUT("/history-village", sejarahHandler.UpdateSejarah)
		protected.DELETE("/history-village", sejarahHandler.DeleteSejarah)

		// Visi Misi
		protected.GET("/visi-misi", visiMisiHandler.GetVisiMisi)
		protected.POST("/visi-misi", visiMisiHandler.CreateVisiMisi)
		protected.PUT("/visi-misi", visiMisiHandler.UpdateVisiMisi)
		protected.DELETE("/visi-misi", visiMisiHandler.DeleteVisiMisi)

		// News
		protected.GET("/news", beritaHandler.GetBerita)
		protected.GET("/news/:id", beritaHandler.GetBeritaByID)
		protected.GET("/news/slug/:slug", beritaHandler.GetBeritaBySlug)
		protected.POST("/news", beritaHandler.CreateBerita)
		protected.PUT("/news/:id", beritaHandler.UpdateBerita)
		protected.DELETE("/news/:id", beritaHandler.DeleteBerita)
		protected.GET("/news/category/count", beritaHandler.GetNewsByCategory)

		// UMKM
		protected.GET("/umkm", umkmHandler.GetAllUmkm)
		protected.GET("/umkm/:id", umkmHandler.GetUmkmByID)
		protected.GET("/umkm/slug/:slug", umkmHandler.GetUmkmBySlug)
		protected.GET("/umkm/count-status", umkmHandler.GetCountStatus)
		protected.POST("/umkm", umkmHandler.CreateUmkm)
		protected.PUT("/umkm/:id", umkmHandler.UpdateUmkm)
		protected.DELETE("/umkm/:id", umkmHandler.DeleteUmkm)

		// Layanan
		protected.GET("/services", layananHandler.GetAllLayanan)
		protected.GET("/service/:id", layananHandler.GetLayananByID)
		protected.GET("/service/slug/:slug", layananHandler.GetLayananBySlug)
		protected.POST("/service", layananHandler.CreateLayanan)
		protected.PUT("/service/:id", layananHandler.UpdateLayanan)
		protected.DELETE("/service/:id", layananHandler.DeleteLayanan)

		// Sub-Layanan
		protected.GET("/sub-service", subLayananHandler.GetAllSubLayanan)
		protected.GET("/sub-service/:id", subLayananHandler.GetSubLayananByID)
		protected.POST("/sub-service", subLayananHandler.CreateSubLayanan)
		protected.PUT("/sub-service/:id", subLayananHandler.UpdateSubLayanan)
		protected.DELETE("/sub-service/:id", subLayananHandler.DeleteSubLayanan)

		// Pengajuan
		protected.GET("/sub-service/submissions", pengajuanHandler.GetAllPengajuan)
		protected.POST("/sub-service/submission", pengajuanHandler.CreatePengajuan)
		protected.GET("/sub-service/submission/:id", pengajuanHandler.GetPengajuanByID)
		protected.GET("/sub-service/submission/nik/:nik", pengajuanHandler.GetPengajuanByNIK)
		protected.GET("/sub-service/submission/surat/:nomor_surat", pengajuanHandler.GetPengajuanByNomorSurat)
		protected.PUT("/sub-service/submission/:id", pengajuanHandler.UpdatePengajuan)
		protected.PUT("/sub-service/submission/status/:id", pengajuanHandler.UpdateStatusPengajuan)
		protected.PUT("/sub-service/submission/approve/:id", pengajuanHandler.ApprovePengajuan)
		protected.PUT("/sub-service/submission/reject/:id", pengajuanHandler.RejectPengajuan)
		protected.DELETE("/sub-service/submission/:id", pengajuanHandler.DeletePengajuan)
		protected.GET("/sub-service/statistics", pengajuanHandler.GetStatisticsByStatus)
		protected.GET("/sub-service/history/all/:pengajuan_id", riwayatHandler.GetAllRiwayat)
		protected.GET("/sub-service/history/:pengajuan_id", riwayatHandler.GetAllRiwayat)

		// APBD
		protected.GET("/apbd", apbdHandler.GetApbd)
		protected.GET("/apbd/:id", apbdHandler.GetApbdByID)
		protected.POST("/apbd", apbdHandler.CreateApbd)
		protected.PUT("/apbd/:id", apbdHandler.UpdateApbd)
		protected.DELETE("/apbd/:id", apbdHandler.DeleteApbd)

		// Penduduk
		protected.GET("/residents", pendudukHandler.GetAllPenduduk)
		protected.GET("/resident/:id", pendudukHandler.GetPendudukByID)
		protected.POST("/resident", pendudukHandler.CreatePenduduk)
		protected.GET("/resident/count", pendudukHandler.CountPenduduk)
		protected.PUT("/resident/:id", pendudukHandler.UpdatePenduduk)
		protected.DELETE("/resident/:id", pendudukHandler.DeletePenduduk)

		// Pendidikan
		protected.GET("/education/program", pendidikanHandler.GetAllProgram)
		protected.GET("/education/program/:id", pendidikanHandler.GetProgramByID)
		protected.POST("/education/program", pendidikanHandler.CreateProgram)
		protected.PUT("/education/program/:id", pendidikanHandler.UpdateProgram)
		protected.DELETE("/education/program/:id", pendidikanHandler.DeleteProgram)
		protected.GET("/education/institution", pendidikanHandler.GetAllLembaga)
		protected.GET("/education/institution/:id", pendidikanHandler.GetLembagaByID)
		protected.POST("/education/institution", pendidikanHandler.CreateLembaga)
		protected.PUT("/education/institution/:id", pendidikanHandler.UpdateLembaga)
		protected.DELETE("/education/institution/:id", pendidikanHandler.DeleteLembaga)
		protected.GET("/education/statistic", pendidikanHandler.GetAllStatistik)
		protected.GET("/education/statistic/:id", pendidikanHandler.GetStatistikByID)
		protected.POST("/education/statistic", pendidikanHandler.CreateStatistik)
		protected.PUT("/education/statistic/:id", pendidikanHandler.UpdateStatistik)
		protected.DELETE("/education/statistic/:id", pendidikanHandler.DeleteStatistik)
		protected.GET("/education/achievements", pendidikanHandler.GetAllCapaian)
		protected.GET("/education/achievements/:id", pendidikanHandler.GetCapaianByID)
		protected.POST("/education/achievements", pendidikanHandler.CreateCapaian)
		protected.PUT("/education/achievements/:id", pendidikanHandler.UpdateCapaian)
		protected.DELETE("/education/achievements/:id", pendidikanHandler.DeleteCapaian)
		protected.GET("/education/documentation", pendidikanHandler.GetAllDokumentasi)
		protected.GET("/education/documentation/:id", pendidikanHandler.GetDokumentasiByID)
		protected.POST("/education/documentation", pendidikanHandler.CreateDokumentasi)
		protected.PUT("/education/documentation/:id", pendidikanHandler.UpdateDokumentasi)
		protected.DELETE("/education/documentation/:id", pendidikanHandler.DeleteDokumentasi)

		// Kesehatan
		protected.GET("/health/service", kesehatanHandler.GetLayanan)
		protected.GET("/health/service/:id", kesehatanHandler.GetLayananByID)
		protected.POST("/health/service", kesehatanHandler.CreateLayanan)
		protected.PUT("/health/service/:id", kesehatanHandler.UpdateLayanan)
		protected.DELETE("/health/service/:id", kesehatanHandler.DeleteLayanan)
		protected.GET("/health/facility", kesehatanHandler.GetFasilitasKesehatan)
		protected.GET("/health/facility/:id", kesehatanHandler.GetFasilitasKesehatanByID)
		protected.POST("/health/facility", kesehatanHandler.CreateFasilitas)
		protected.PUT("/health/facility/:id", kesehatanHandler.UpdateFasilitasKesehatan)
		protected.DELETE("/health/facility/:id", kesehatanHandler.DeleteFasilitasKesehatan)

		// Dashboard
		protected.GET("/dashboard/summary", dashboardHandler.GetDashboardSummary)
	}
}
