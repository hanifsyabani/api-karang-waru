package usecase

import (
	"api-karang-waru/internal/domain"
	"gorm.io/gorm"
)

type KesehatanRepository interface {
	CreateLayanan(layanan *domain.LayananKesehatan) error
	FindAllLayanan() ([]domain.LayananKesehatan, error)
	FindLayananByID(id uint) (*domain.LayananKesehatan, error)
	UpdateLayanan(layanan *domain.LayananKesehatan) error
	DeleteLayanan(layanan *domain.LayananKesehatan) error

	CreateFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error
	FindAllFasilitasKesehatan() ([]domain.FasilitasKesehatan, error)
	FindFasilitasKesehatanByID(id uint) (*domain.FasilitasKesehatan, error)
	UpdateFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error
	DeleteFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error
}

type kesehatanRepository struct {
	db *gorm.DB
}

func NewKesehatanRepository(db *gorm.DB) KesehatanRepository {
	return &kesehatanRepository{db}
}

func (r *kesehatanRepository) CreateLayanan(layanan *domain.LayananKesehatan) error {
	return r.db.Create(layanan).Error
}

func (r *kesehatanRepository) FindAllLayanan() ([]domain.LayananKesehatan, error) {
	var layanan []domain.LayananKesehatan
	err := r.db.Preload("Fasilitas").Find(&layanan).Error
	return layanan, err
}

func (r *kesehatanRepository) FindLayananByID(id uint) (*domain.LayananKesehatan, error) {
	var layanan domain.LayananKesehatan
	err := r.db.Preload("Fasilitas").First(&layanan, id).Error
	return &layanan, err
}

func (r *kesehatanRepository) UpdateLayanan(layanan *domain.LayananKesehatan) error {
	return r.db.Save(layanan).Error
}

func (r *kesehatanRepository) DeleteLayanan(layanan *domain.LayananKesehatan) error {
	return r.db.Delete(layanan).Error
}


func (r *kesehatanRepository) CreateFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error {
	return r.db.Create(fasilitas).Error
}

func (r *kesehatanRepository) FindAllFasilitasKesehatan() ([]domain.FasilitasKesehatan, error) {
	var fasilitas []domain.FasilitasKesehatan
	err := r.db.Find(&fasilitas).Error
	return fasilitas, err
}

func (r *kesehatanRepository) FindFasilitasKesehatanByID(id uint) (*domain.FasilitasKesehatan, error) {
	var fasilitas domain.FasilitasKesehatan
	err := r.db.Find(&fasilitas, id).Error
	return &fasilitas, err
}

func (r *kesehatanRepository) UpdateFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error {
	return r.db.Save(fasilitas).Error
}

func (r *kesehatanRepository) DeleteFasilitasKesehatan(fasilitas *domain.FasilitasKesehatan) error {
	return r.db.Delete(fasilitas).Error
}

