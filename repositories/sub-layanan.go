package repositories

import (
	"api-karang-waru/models"
	"strings"

	"gorm.io/gorm"
)

type SubLayananRepository interface {
	CreateSubLayanan(subLayanan *models.SubLayananDesa) error
	FindSubLayanan(
		layananDesaID uint,
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.SubLayananDesa, error)
	FindSubLayananByID(id uint) (*models.SubLayananDesa, error)
	FindSubLayananByLayananID(layananDesaID uint) ([]models.SubLayananDesa, error)
	UpdateSubLayanan(subLayanan *models.SubLayananDesa) error
	DeleteSubLayanan(subLayanan *models.SubLayananDesa) error
}

type subLayananRepository struct {
	db *gorm.DB
}

func NewSubLayananRepository(db *gorm.DB) SubLayananRepository {
	return &subLayananRepository{db}
}

func (r *subLayananRepository) CreateSubLayanan(subLayanan *models.SubLayananDesa) error {
	return r.db.Create(subLayanan).Error
}

func (r *subLayananRepository) FindSubLayanan(
	layananDesaID uint,
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.SubLayananDesa, error) {
	var subLayanan []models.SubLayananDesa

	offset := (page - 1) * limit
	query := r.db.Model(&models.SubLayananDesa{})

	// Filter by layanan desa ID if provided
	if layananDesaID > 0 {
		query = query.Where("layanan_desa_id = ?", layananDesaID)
	}

	// Search functionality
	if search != "" {
		search = strings.ToLower(search)
		searchPattern := "%" + search + "%"
		query = query.Where("LOWER(nama) LIKE ? OR LOWER(persyaratan) LIKE ?", searchPattern, searchPattern)
	}

	// Validate sort order
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	query = query.Order(sortBy + " " + sortOrder)

	err := query.Offset(offset).Limit(limit).Find(&subLayanan).Error
	return subLayanan, err
}

func (r *subLayananRepository) FindSubLayananByID(id uint) (*models.SubLayananDesa, error) {
	var subLayanan models.SubLayananDesa
	err := r.db.First(&subLayanan, id).Error
	return &subLayanan, err
}

func (r *subLayananRepository) FindSubLayananByLayananID(layananDesaID uint) ([]models.SubLayananDesa, error) {
	var subLayanan []models.SubLayananDesa
	err := r.db.Where("layanan_desa_id = ? AND aktif = ?", layananDesaID, true).Find(&subLayanan).Error
	return subLayanan, err
}

func (r *subLayananRepository) UpdateSubLayanan(subLayanan *models.SubLayananDesa) error {
	return r.db.Save(subLayanan).Error
}

func (r *subLayananRepository) DeleteSubLayanan(subLayanan *models.SubLayananDesa) error {
	return r.db.Delete(subLayanan).Error
}


type PengajuanLayananRepository interface {
	CreatePengajuan(pengajuan *models.PengajuanLayanan) error
	FindPengajuan(
		search string,
		status string,
		nik string,
		layananDesaID uint,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.PengajuanLayanan, error)
	FindPengajuanByID(id uint) (*models.PengajuanLayanan, error)
	FindPengajuanByNIK(nik string) ([]models.PengajuanLayanan, error)
	FindPengajuanByNomorSurat(nomorSurat string) (*models.PengajuanLayanan, error)
	UpdatePengajuan(pengajuan *models.PengajuanLayanan) error
	DeletePengajuan(pengajuan *models.PengajuanLayanan) error
	CountByStatus(status string) (int64, error)
}

type pengajuanLayananRepository struct {
	db *gorm.DB
}

func NewPengajuanLayananRepository(db *gorm.DB) PengajuanLayananRepository {
	return &pengajuanLayananRepository{db}
}

func (r *pengajuanLayananRepository) CreatePengajuan(pengajuan *models.PengajuanLayanan) error {
	return r.db.Create(pengajuan).Error
}

func (r *pengajuanLayananRepository) FindPengajuan(
	search string,
	status string,
	nik string,
	layananDesaID uint,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.PengajuanLayanan, error) {
	var pengajuan []models.PengajuanLayanan

	offset := (page - 1) * limit
	query := r.db.Model(&models.PengajuanLayanan{}).
		Preload("LayananDesa").
		Preload("SubLayanan")

	// Filter by status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by NIK
	if nik != "" {
		query = query.Where("nik = ?", nik)
	}

	// Filter by layanan desa ID
	if layananDesaID > 0 {
		query = query.Where("layanan_desa_id = ?", layananDesaID)
	}

	// Search functionality
	if search != "" {
		search = strings.ToLower(search)
		searchPattern := "%" + search + "%"
		query = query.Where(
			"LOWER(nama_lengkap) LIKE ? OR LOWER(nik) LIKE ? OR LOWER(nomor_surat) LIKE ? OR LOWER(keperluan) LIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// Validate sort order
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	query = query.Order(sortBy + " " + sortOrder)

	err := query.Offset(offset).Limit(limit).Find(&pengajuan).Error
	return pengajuan, err
}

func (r *pengajuanLayananRepository) FindPengajuanByID(id uint) (*models.PengajuanLayanan, error) {
	var pengajuan models.PengajuanLayanan
	err := r.db.Preload("LayananDesa").
		Preload("SubLayanan").
		Preload("Riwayat").
		First(&pengajuan, id).Error
	return &pengajuan, err
}

func (r *pengajuanLayananRepository) FindPengajuanByNIK(nik string) ([]models.PengajuanLayanan, error) {
	var pengajuan []models.PengajuanLayanan
	err := r.db.Where("nik = ?", nik).
		Preload("LayananDesa").
		Preload("SubLayanan").
		Order("created_at DESC").
		Find(&pengajuan).Error
	return pengajuan, err
}

func (r *pengajuanLayananRepository) FindPengajuanByNomorSurat(nomorSurat string) (*models.PengajuanLayanan, error) {
	var pengajuan models.PengajuanLayanan
	err := r.db.Where("nomor_surat = ?", nomorSurat).
		Preload("LayananDesa").
		Preload("SubLayanan").
		First(&pengajuan).Error
	return &pengajuan, err
}

func (r *pengajuanLayananRepository) UpdatePengajuan(pengajuan *models.PengajuanLayanan) error {
	return r.db.Save(pengajuan).Error
}

func (r *pengajuanLayananRepository) DeletePengajuan(pengajuan *models.PengajuanLayanan) error {
	return r.db.Delete(pengajuan).Error
}

func (r *pengajuanLayananRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&models.PengajuanLayanan{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

// ==================== RIWAYAT PENGAJUAN REPOSITORY ====================

type RiwayatPengajuanRepository interface {
	CreateRiwayat(riwayat *models.RiwayatPengajuan) error
	FindRiwayatByPengajuanID(pengajuanID uint) ([]models.RiwayatPengajuan, error)
	FindRiwayat(
		pengajuanID uint,
		page int,
		limit int,
	) ([]models.RiwayatPengajuan, error)
}

type riwayatPengajuanRepository struct {
	db *gorm.DB
}

func NewRiwayatPengajuanRepository(db *gorm.DB) RiwayatPengajuanRepository {
	return &riwayatPengajuanRepository{db}
}

func (r *riwayatPengajuanRepository) CreateRiwayat(riwayat *models.RiwayatPengajuan) error {
	return r.db.Create(riwayat).Error
}

func (r *riwayatPengajuanRepository) FindRiwayatByPengajuanID(pengajuanID uint) ([]models.RiwayatPengajuan, error) {
	var riwayat []models.RiwayatPengajuan
	err := r.db.Where("pengajuan_id = ?", pengajuanID).
		Order("created_at DESC").
		Find(&riwayat).Error
	return riwayat, err
}

func (r *riwayatPengajuanRepository) FindRiwayat(
	pengajuanID uint,
	page int,
	limit int,
) ([]models.RiwayatPengajuan, error) {
	var riwayat []models.RiwayatPengajuan

	offset := (page - 1) * limit
	query := r.db.Model(&models.RiwayatPengajuan{})

	if pengajuanID > 0 {
		query = query.Where("pengajuan_id = ?", pengajuanID)
	}

	err := query.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&riwayat).Error
	return riwayat, err
}