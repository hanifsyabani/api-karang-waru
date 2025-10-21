package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type LayananRepository interface {
	CreateLayanan(layanan *models.LayananDesa) error
	FindLayanan() ([]models.LayananDesa, error)
	FindLayananByID(id uint) (*models.LayananDesa, error)
	FindLayananBySlug(slug string) (*models.LayananDesa, error)
	UpdateLayanan(layanan *models.LayananDesa) error
	DeleteLayanan(layanan *models.LayananDesa) error
}

type layananRepository struct {
	db *gorm.DB
}

func NewLayananRepository(db *gorm.DB) LayananRepository {
	return &layananRepository{db}
}

func (r *layananRepository) CreateLayanan(layanan *models.LayananDesa) error {
	return r.db.Create(layanan).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *layananRepository) FindLayanan() ([]models.LayananDesa, error) {
	var layanan []models.LayananDesa
	err := r.db.Find(&layanan).Error
	return layanan, err
}

func (r *layananRepository) FindLayananByID(id uint) (*models.LayananDesa, error) {
	var layanan models.LayananDesa
	err := r.db.Find(&layanan, id).Error
	return &layanan, err
}

func (r *layananRepository) FindLayananBySlug(slug string) (*models.LayananDesa, error) {
	var layanan models.LayananDesa
	err := r.db.Where("slug = ?", slug).First(&layanan).Error
	return &layanan, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *layananRepository) UpdateLayanan(layanan *models.LayananDesa) error {
	return r.db.Save(layanan).Error
}

func (r *layananRepository) DeleteLayanan(layanan *models.LayananDesa) error {
	return r.db.Delete(layanan).Error

}
