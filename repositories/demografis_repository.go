package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type DemografisRepository interface {
	CreateDemografis(demografis *models.Demografis) error
	FindDemografis() (models.Demografis, error)
	UpdateDemografis(demografis *models.Demografis) error
	DeleteDemografis() error
}

type demografisRepository struct {
	db *gorm.DB
}

func NewDemografisRepository(db *gorm.DB) DemografisRepository {
	return &demografisRepository{db}
}

func (r *demografisRepository) CreateDemografis(profil *models.Demografis) error {
	return r.db.Create(profil).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *demografisRepository) FindDemografis() (models.Demografis, error) {
	var profil models.Demografis
	err := r.db.First(&profil).Error
	return profil, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *demografisRepository) UpdateDemografis(profil *models.Demografis) error {
	return r.db.Save(profil).Error
}

func (r *demografisRepository) DeleteDemografis() error {
	return r.db.Exec("DELETE FROM demografis").Error
}