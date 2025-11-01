package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type ApbdDesaRepository interface {
	CreateApbd(apbd *models.APBDDesa) error
	FindApbd() ([]models.APBDDesa, error)
	FindApbdByID(id uint) (*models.APBDDesa, error)
	UpdateApbd(apbd *models.APBDDesa) error
	DeleteApbd(apbd *models.APBDDesa) error
}

type apbdRepository struct {
	db *gorm.DB
}

func NewApbdRepository(db *gorm.DB) ApbdDesaRepository {
	return &apbdRepository{db}
}

func (r *apbdRepository) CreateApbd(apbd *models.APBDDesa) error {
	return r.db.Create(apbd).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *apbdRepository) FindApbd() ([]models.APBDDesa, error) {
	var apbd []models.APBDDesa
	err := r.db.Find(&apbd).Error
	return apbd, err
}

func (r *apbdRepository) FindApbdByID(id uint) (*models.APBDDesa, error) {
	var apbd models.APBDDesa
	err := r.db.Find(&apbd, id).Error
	return &apbd, err
}


// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *apbdRepository) UpdateApbd(apbd *models.APBDDesa) error {
	return r.db.Save(apbd).Error
}

func (r *apbdRepository) DeleteApbd(apbd *models.APBDDesa) error {
	return r.db.Delete(apbd).Error

}
