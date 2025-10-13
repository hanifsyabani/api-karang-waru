package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type UmkmRepository interface {
	CreateUmkm(umkm *models.Umkm) error
	FindUmkm() ([]models.Umkm, error)
	FindUmkmByID(id uint) (*models.Umkm, error)
	FindUmkmBySlug(slug string) (*models.Umkm, error)
	UpdateUmkm(umkm *models.Umkm) error
	DeleteUmkm(umkm *models.Umkm) error
}

type umkmRepository struct {
	db *gorm.DB
}

func NewUmkmRepository(db *gorm.DB) UmkmRepository {
	return &umkmRepository{db}
}

func (r *umkmRepository) CreateUmkm(umkm *models.Umkm) error {
	return r.db.Create(umkm).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *umkmRepository) FindUmkm() ([]models.Umkm, error) {
	var umkm []models.Umkm
	err := r.db.Find(&umkm).Error
	return umkm, err
}

func (r *umkmRepository) FindUmkmByID(id uint) (*models.Umkm, error) {
	var umkm models.Umkm
	err := r.db.Find(&umkm, id).Error
	return &umkm, err
}

func (r *umkmRepository) FindUmkmBySlug(slug string) (*models.Umkm, error) {
	var umkm models.Umkm
	err := r.db.Where("slug = ?", slug).First(&umkm).Error
	return &umkm, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *umkmRepository) UpdateUmkm(umkm *models.Umkm) error {
	return r.db.Save(umkm).Error
}

func (r *umkmRepository) DeleteUmkm(umkm *models.Umkm) error {
	return r.db.Delete(umkm).Error

}
