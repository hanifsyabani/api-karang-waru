package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type PendudukRepository interface {
	CreatePenduduk(penduduk *models.Penduduk) error
	FindPenduduk() ([]models.Penduduk, error)
	FindPendudukByID(id uint) (*models.Penduduk, error)
	UpdatePenduduk(penduduk *models.Penduduk) error
	DeletePenduduk(penduduk *models.Penduduk) error
}

type pendudukRepository struct {
	db *gorm.DB
}

func NewPendudukRepository(db *gorm.DB) PendudukRepository {
	return &pendudukRepository{db}
}

func (r *pendudukRepository) CreatePenduduk(penduduk *models.Penduduk) error {
	return r.db.Create(penduduk).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *pendudukRepository) FindPenduduk() ([]models.Penduduk, error) {
	var penduduk []models.Penduduk
	err := r.db.Find(&penduduk).Error
	return penduduk, err
}

func (r *pendudukRepository) FindPendudukByID(id uint) (*models.Penduduk, error) {
	var penduduk models.Penduduk
	err := r.db.Find(&penduduk, id).Error
	return &penduduk, err
}


// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *pendudukRepository) UpdatePenduduk(penduduk *models.Penduduk) error {
	return r.db.Save(penduduk).Error
}

func (r *pendudukRepository) DeletePenduduk(penduduk *models.Penduduk) error {
	return r.db.Delete(penduduk).Error

}
