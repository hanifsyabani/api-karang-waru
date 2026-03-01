package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type ProfilRepository interface {
	CreateProfil(profil *domain.ProfilDesa) error
	FindProfil() (domain.ProfilDesa, error)
	UpdateProfil(profil *domain.ProfilDesa) error
	DeleteProfil() error
}

type profilRepository struct {
	db *gorm.DB
}

func NewProfilRepository(db *gorm.DB) ProfilRepository {
	return &profilRepository{db}
}

func (r *profilRepository) CreateProfil(profil *domain.ProfilDesa) error {
	return r.db.Create(profil).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *profilRepository) FindProfil() (domain.ProfilDesa, error) {
	var profil domain.ProfilDesa
	err := r.db.First(&profil).Error
	return profil, err
}


// tidak butuh id karena langsung method Save() cari priamry key di struct domain.Demografis
func (r *profilRepository) UpdateProfil(profil *domain.ProfilDesa) error {
	return r.db.Save(profil).Error
}

func (r *profilRepository) DeleteProfil() error {
	return r.db.Exec("DELETE FROM profil_desa").Error
}



