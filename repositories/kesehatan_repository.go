package repositories

import (
	"api-karang-waru/models"
	"gorm.io/gorm"
)

type KesehatanRepository interface {
	CreateLayanan(layanan *models.LayananKesehatan) error
	FindAllLayanan() ([]models.LayananKesehatan, error)
	FindLayananByID(id uint) (*models.LayananKesehatan, error)
	UpdateLayanan(layanan *models.LayananKesehatan) error
	DeleteLayanan(layanan *models.LayananKesehatan) error

	CreateFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error
	FindAllFasilitasKesehatan() ([]models.FasilitasKesehatan, error)
	FindFasilitasKesehatanByID(id uint) (*models.FasilitasKesehatan, error)
	UpdateFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error
	DeleteFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error
}

type kesehatanRepository struct {
	db *gorm.DB
}

func NewKesehatanRepository(db *gorm.DB) KesehatanRepository {
	return &kesehatanRepository{db}
}

func (r *kesehatanRepository) CreateLayanan(layanan *models.LayananKesehatan) error {
	return r.db.Create(layanan).Error
}

func (r *kesehatanRepository) FindAllLayanan() ([]models.LayananKesehatan, error) {
	var layanan []models.LayananKesehatan
	err := r.db.Preload("Fasilitas").Find(&layanan).Error
	return layanan, err
}

func (r *kesehatanRepository) FindLayananByID(id uint) (*models.LayananKesehatan, error) {
	var layanan models.LayananKesehatan
	err := r.db.Preload("Fasilitas").First(&layanan, id).Error
	return &layanan, err
}

func (r *kesehatanRepository) UpdateLayanan(layanan *models.LayananKesehatan) error {
	return r.db.Save(layanan).Error
}

func (r *kesehatanRepository) DeleteLayanan(layanan *models.LayananKesehatan) error {
	return r.db.Delete(layanan).Error
}


func (r *kesehatanRepository) CreateFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error {
	return r.db.Create(fasilitas).Error
}

func (r *kesehatanRepository) FindAllFasilitasKesehatan() ([]models.FasilitasKesehatan, error) {
	var fasilitas []models.FasilitasKesehatan
	err := r.db.Find(&fasilitas).Error
	return fasilitas, err
}

func (r *kesehatanRepository) FindFasilitasKesehatanByID(id uint) (*models.FasilitasKesehatan, error) {
	var fasilitas models.FasilitasKesehatan
	err := r.db.Find(&fasilitas, id).Error
	return &fasilitas, err
}

func (r *kesehatanRepository) UpdateFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error {
	return r.db.Save(fasilitas).Error
}

func (r *kesehatanRepository) DeleteFasilitasKesehatan(fasilitas *models.FasilitasKesehatan) error {
	return r.db.Delete(fasilitas).Error
}
