package usecase

import (
	"api-karang-waru/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type LayananRepository interface {
	CreateLayanan(layanan *domain.LayananDesa) error
	FindLayanan(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.LayananDesa, error)
	FindLayananByID(id uint) (*domain.LayananDesa, error)
	FindLayananBySlug(slug string) (*domain.LayananDesa, error)
	UpdateLayanan(layanan *domain.LayananDesa) error
	DeleteLayanan(layanan *domain.LayananDesa) error
}

type layananRepository struct {
	db *gorm.DB
}

func NewLayananRepository(db *gorm.DB) LayananRepository {
	return &layananRepository{db}
}

func (r *layananRepository) CreateLayanan(layanan *domain.LayananDesa) error {
	return r.db.Create(layanan).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *layananRepository) FindLayanan(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.LayananDesa, error) {
	var layanan []domain.LayananDesa

	offset := (page - 1) * limit
	query := r.db.Model(&domain.LayananDesa{})

	if search != "" {
		search = strings.ToLower(search)
		searchPattern := "%" + search + "%"
		query = query.Where("LOWER(service_name) LIKE ? OR LOWER(description) LIKE ?", searchPattern, searchPattern)
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	query = query.Order(sortBy + " " + sortOrder)

	err := query.Offset(offset).Limit(limit).Find(&layanan).Error
	return layanan, err
}

func (r *layananRepository) FindLayananByID(id uint) (*domain.LayananDesa, error) {
	var layanan domain.LayananDesa
	err := r.db.Find(&layanan, id).Error
	return &layanan, err
}

func (r *layananRepository) FindLayananBySlug(slug string) (*domain.LayananDesa, error) {
	var layanan domain.LayananDesa
	err := r.db.Where("slug = ?", slug).First(&layanan).Error
	return &layanan, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *layananRepository) UpdateLayanan(layanan *domain.LayananDesa) error {
	return r.db.Save(layanan).Error
}

func (r *layananRepository) DeleteLayanan(layanan *domain.LayananDesa) error {
	return r.db.Delete(layanan).Error

}

