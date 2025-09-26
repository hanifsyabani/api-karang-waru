package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type ProfilRepository interface {
	CreateProfil(profil *models.ProfilDesa) error
	FindProfil() (models.ProfilDesa, error)
	UpdateProfil(profil *models.ProfilDesa) error
	DeleteProfil() error
}

type profilRepository struct {
	db *gorm.DB
}

func NewProfilRepository(db *gorm.DB) ProfilRepository {
	return &profilRepository{db}
}

func (r *profilRepository) CreateProfil(profil *models.ProfilDesa) error {
	return r.db.Create(profil).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *profilRepository) FindProfil() (models.ProfilDesa, error) {
	var profil models.ProfilDesa
	err := r.db.First(&profil).Error
	return profil, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.Demografis
func (r *profilRepository) UpdateProfil(profil *models.ProfilDesa) error {
	return r.db.Save(profil).Error
}

func (r *profilRepository) DeleteProfil() error {
	return r.db.Exec("DELETE FROM profil_desa").Error
}


