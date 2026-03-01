package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type DemografisRepository interface {
	CreateDemografis(demografis *domain.Demografis) error
	FindDemografis() (domain.Demografis, error)
	UpdateDemografis(demografis *domain.Demografis) error
	DeleteDemografis() error
}

type demografisRepository struct {
	db *gorm.DB
}

func NewDemografisRepository(db *gorm.DB) DemografisRepository {
	return &demografisRepository{db}
}

func (r *demografisRepository) CreateDemografis(profil *domain.Demografis) error {
	return r.db.Create(profil).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *demografisRepository) FindDemografis() (domain.Demografis, error) {
	var profil domain.Demografis
	err := r.db.First(&profil).Error
	return profil, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *demografisRepository) UpdateDemografis(profil *domain.Demografis) error {
	return r.db.Save(profil).Error
}

func (r *demografisRepository) DeleteDemografis() error {
	return r.db.Exec("DELETE FROM demografis").Error
}
