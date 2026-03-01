package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type SejarahRepository interface {
	CreateSejarah(sejarah *domain.Sejarah) error
	FindSejarah() (domain.Sejarah, error)
	UpdateSejarah(sejarah *domain.Sejarah) error
	DeleteSejarah() error
}

type sejarahRepository struct {
	db *gorm.DB
}

func NewSejarahRepository(db *gorm.DB) SejarahRepository {
	return &sejarahRepository{db}
}

func (r *sejarahRepository) CreateSejarah(sejarah *domain.Sejarah) error {
	return r.db.Create(sejarah).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *sejarahRepository) FindSejarah() (domain.Sejarah, error) {
	var sejarah domain.Sejarah
	err := r.db.First(&sejarah).Error
	return sejarah, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *sejarahRepository) UpdateSejarah(sejarah *domain.Sejarah) error {
	return r.db.Save(sejarah).Error
}

func (r *sejarahRepository) DeleteSejarah() error {
	return r.db.Exec("DELETE FROM sejarah").Error
}
