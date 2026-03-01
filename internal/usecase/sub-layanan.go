package usecase

import (
	"api-karang-waru/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type SubLayananRepository interface {
	CreateSubLayanan(subLayanan *domain.SubLayananDesa) error
	FindSubLayanan(
		layananDesaID uint,
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.SubLayananDesa, error)
	FindSubLayananByID(id uint) (*domain.SubLayananDesa, error)
	FindSubLayananByLayananID(layananDesaID uint) ([]domain.SubLayananDesa, error)
	UpdateSubLayanan(subLayanan *domain.SubLayananDesa) error
	DeleteSubLayanan(subLayanan *domain.SubLayananDesa) error
}

type subLayananRepository struct {
	db *gorm.DB
}

func NewSubLayananRepository(db *gorm.DB) SubLayananRepository {
	return &subLayananRepository{db}
}

func (r *subLayananRepository) CreateSubLayanan(subLayanan *domain.SubLayananDesa) error {
	return r.db.Create(subLayanan).Error
}

func (r *subLayananRepository) FindSubLayanan(
	layananDesaID uint,
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.SubLayananDesa, error) {
	var subLayanan []domain.SubLayananDesa

	offset := (page - 1) * limit
	query := r.db.Model(&domain.SubLayananDesa{})

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

func (r *subLayananRepository) FindSubLayananByID(id uint) (*domain.SubLayananDesa, error) {
	var subLayanan domain.SubLayananDesa
	err := r.db.First(&subLayanan, id).Error
	return &subLayanan, err
}

func (r *subLayananRepository) FindSubLayananByLayananID(layananDesaID uint) ([]domain.SubLayananDesa, error) {
	var subLayanan []domain.SubLayananDesa
	err := r.db.Where("layanan_desa_id = ? AND aktif = ?", layananDesaID, true).Find(&subLayanan).Error
	return subLayanan, err
}

func (r *subLayananRepository) UpdateSubLayanan(subLayanan *domain.SubLayananDesa) error {
	return r.db.Save(subLayanan).Error
}

func (r *subLayananRepository) DeleteSubLayanan(subLayanan *domain.SubLayananDesa) error {
	return r.db.Delete(subLayanan).Error
}


type PengajuanLayananRepository interface {
	CreatePengajuan(pengajuan *domain.PengajuanLayanan) error
	FindPengajuan(
		search string,
		status string,
		nik string,
		layananDesaID uint,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.PengajuanLayanan, error)
	FindPengajuanByID(id uint) (*domain.PengajuanLayanan, error)
	FindPengajuanByNIK(nik string) ([]domain.PengajuanLayanan, error)
	FindPengajuanByNomorSurat(nomorSurat string) (*domain.PengajuanLayanan, error)
	UpdatePengajuan(pengajuan *domain.PengajuanLayanan) error
	DeletePengajuan(pengajuan *domain.PengajuanLayanan) error
	CountByStatus(status string) (int64, error)
}

type pengajuanLayananRepository struct {
	db *gorm.DB
}

func NewPengajuanLayananRepository(db *gorm.DB) PengajuanLayananRepository {
	return &pengajuanLayananRepository{db}
}

func (r *pengajuanLayananRepository) CreatePengajuan(pengajuan *domain.PengajuanLayanan) error {
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
) ([]domain.PengajuanLayanan, error) {
	var pengajuan []domain.PengajuanLayanan

	offset := (page - 1) * limit
	query := r.db.Model(&domain.PengajuanLayanan{}).
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

func (r *pengajuanLayananRepository) FindPengajuanByID(id uint) (*domain.PengajuanLayanan, error) {
	var pengajuan domain.PengajuanLayanan
	err := r.db.Preload("LayananDesa").
		Preload("SubLayanan").
		Preload("Riwayat").
		First(&pengajuan, id).Error
	return &pengajuan, err
}

func (r *pengajuanLayananRepository) FindPengajuanByNIK(nik string) ([]domain.PengajuanLayanan, error) {
	var pengajuan []domain.PengajuanLayanan
	err := r.db.Where("nik = ?", nik).
		Preload("LayananDesa").
		Preload("SubLayanan").
		Order("created_at DESC").
		Find(&pengajuan).Error
	return pengajuan, err
}

func (r *pengajuanLayananRepository) FindPengajuanByNomorSurat(nomorSurat string) (*domain.PengajuanLayanan, error) {
	var pengajuan domain.PengajuanLayanan
	err := r.db.Where("nomor_surat = ?", nomorSurat).
		Preload("LayananDesa").
		Preload("SubLayanan").
		First(&pengajuan).Error
	return &pengajuan, err
}

func (r *pengajuanLayananRepository) UpdatePengajuan(pengajuan *domain.PengajuanLayanan) error {
	return r.db.Save(pengajuan).Error
}

func (r *pengajuanLayananRepository) DeletePengajuan(pengajuan *domain.PengajuanLayanan) error {
	return r.db.Delete(pengajuan).Error
}

func (r *pengajuanLayananRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.PengajuanLayanan{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

// ==================== RIWAYAT PENGAJUAN REPOSITORY ====================

type RiwayatPengajuanRepository interface {
	CreateRiwayat(riwayat *domain.RiwayatPengajuan) error
	FindRiwayatByPengajuanID(pengajuanID uint) ([]domain.RiwayatPengajuan, error)
	FindRiwayat(
		pengajuanID uint,
		page int,
		limit int,
	) ([]domain.RiwayatPengajuan, error)
}

type riwayatPengajuanRepository struct {
	db *gorm.DB
}

func NewRiwayatPengajuanRepository(db *gorm.DB) RiwayatPengajuanRepository {
	return &riwayatPengajuanRepository{db}
}

func (r *riwayatPengajuanRepository) CreateRiwayat(riwayat *domain.RiwayatPengajuan) error {
	return r.db.Create(riwayat).Error
}

func (r *riwayatPengajuanRepository) FindRiwayatByPengajuanID(pengajuanID uint) ([]domain.RiwayatPengajuan, error) {
	var riwayat []domain.RiwayatPengajuan
	err := r.db.Where("pengajuan_id = ?", pengajuanID).
		Order("created_at DESC").
		Find(&riwayat).Error
	return riwayat, err
}

func (r *riwayatPengajuanRepository) FindRiwayat(
	pengajuanID uint,
	page int,
	limit int,
) ([]domain.RiwayatPengajuan, error) {
	var riwayat []domain.RiwayatPengajuan

	offset := (page - 1) * limit
	query := r.db.Model(&domain.RiwayatPengajuan{})

	if pengajuanID > 0 {
		query = query.Where("pengajuan_id = ?", pengajuanID)
	}

	err := query.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&riwayat).Error
	return riwayat, err
}
