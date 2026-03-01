package handler

import (
	"api-karang-waru/config"
	"api-karang-waru/internal/domain"
	"api-karang-waru/pkg/types"
	"net/http"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct{}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) GetDashboardSummary(c *gin.Context) {
	db := config.DB
	result := make(map[string]interface{})

	// ---- Penduduk ----
	var totalPenduduk, lakiLaki, perempuan, kartuKeluarga int64
	db.Model(&domain.Penduduk{}).Count(&totalPenduduk)
	db.Model(&domain.Penduduk{}).Where("jenis_kelamin = ?", "Laki-laki").Count(&lakiLaki)
	db.Model(&domain.Penduduk{}).Where("jenis_kelamin = ?", "Perempuan").Count(&perempuan)
	db.Model(&domain.Penduduk{}).Distinct("no_kk").Count(&kartuKeluarga)
	result["penduduk"] = map[string]interface{}{
		"total":             totalPenduduk,
		"male":              lakiLaki,
		"female":            perempuan,
		"family_card_count": kartuKeluarga,
	}

	// ---- Berita ----
	var totalBerita int64
	db.Model(&domain.Berita{}).Count(&totalBerita)

	type CategoryCount struct {
		Category string `json:"category"`
		Count    int64  `json:"count"`
	}
	var beritaByCategory []CategoryCount
	db.Model(&domain.Berita{}).Select("category, COUNT(*) as count").Group("category").Scan(&beritaByCategory)
	result["berita"] = map[string]interface{}{
		"total":       totalBerita,
		"by_category": beritaByCategory,
	}

	// ---- UMKM ----
	var totalUmkm int64
	db.Model(&domain.Umkm{}).Count(&totalUmkm)

	type StatusCount struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}
	var umkmByStatus []StatusCount
	db.Model(&domain.Umkm{}).Select("status, COUNT(*) as count").Group("status").Scan(&umkmByStatus)
	result["umkm"] = map[string]interface{}{
		"total":     totalUmkm,
		"by_status": umkmByStatus,
	}

	// ---- Layanan ----
	var totalLayanan int64
	db.Model(&domain.LayananDesa{}).Count(&totalLayanan)

	var totalPengajuan int64
	db.Model(&domain.PengajuanLayanan{}).Count(&totalPengajuan)

	var pengajuanByStatus []StatusCount
	db.Model(&domain.PengajuanLayanan{}).Select("status, COUNT(*) as count").Group("status").Scan(&pengajuanByStatus)
	result["layanan"] = map[string]interface{}{
		"total_services":    totalLayanan,
		"total_submissions": totalPengajuan,
		"submissions_by_status": pengajuanByStatus,
	}

	// ---- APBD ----
	var apbd domain.APBDDesa
	err := db.Order("tahun DESC").First(&apbd).Error
	if err == nil {
		result["apbd"] = map[string]interface{}{
			"tahun":                          apbd.Tahun,
			"pendapatan_asli_desa":           apbd.PendapatanAsliDesa,
			"transfer":                       apbd.Transfer,
			"pendapatan_lain":                apbd.PendapatanLain,
			"belanja_pemerintahan":           apbd.BelanjaPenyelenggaraanPemerintahan,
			"belanja_pembangunan":            apbd.BelanjaPembangunan,
			"belanja_pembinaan":              apbd.BelanjaPembinaanKemasyarakatan,
			"belanja_pemberdayaan":           apbd.BelanjaPemberdayaanMasyarakat,
			"belanja_takterduga":             apbd.BelanjaTakTerduga,
			"total_pendapatan":               apbd.TotalPendapatan,
			"total_belanja":                  apbd.TotalBelanja,
			"surplus_defisit":                apbd.SurplusDefisit,
		}
	} else {
		result["apbd"] = nil
	}

	// ---- Pendidikan ----
	var statistikPendidikan domain.StatistikPendidikan
	err2 := db.Order("tahun DESC").First(&statistikPendidikan).Error
	if err2 == nil {
		result["pendidikan"] = map[string]interface{}{
			"tahun":            statistikPendidikan.Tahun,
			"tidak_sekolah":    statistikPendidikan.TidakSekolah,
			"sd":               statistikPendidikan.SD,
			"smp":              statistikPendidikan.SMP,
			"sma":              statistikPendidikan.SMA,
			"perguruan_tinggi": statistikPendidikan.PerguruanTinggi,
		}
	} else {
		result["pendidikan"] = nil
	}

	var totalLembaga int64
	db.Model(&domain.LembagaPendidikan{}).Count(&totalLembaga)
	result["total_lembaga_pendidikan"] = totalLembaga

	// ---- Kesehatan ----
	var totalFasilitas, totalLayananKesehatan int64
	db.Model(&domain.FasilitasKesehatan{}).Count(&totalFasilitas)
	db.Model(&domain.LayananKesehatan{}).Count(&totalLayananKesehatan)
	result["kesehatan"] = map[string]interface{}{
		"total_fasilitas": totalFasilitas,
		"total_layanan":   totalLayananKesehatan,
	}

	// ---- Users ----
	var totalUsers int64
	db.Model(&domain.User{}).Count(&totalUsers)
	result["total_users"] = totalUsers

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Dashboard summary retrieved successfully",
		Data:    result,
	})
}

