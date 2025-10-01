package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type SejarahRepository interface {
	CreateSejarah(sejarah *models.Sejarah) error
	FindSejarah() (models.Sejarah, error)
	UpdateSejarah(sejarah *models.Sejarah) error
	DeleteSejarah() error
}

type sejarahRepository struct {
	db *gorm.DB
}

func NewSejarahRepository(db *gorm.DB) SejarahRepository {
	return &sejarahRepository{db}
}

func (r *sejarahRepository) CreateSejarah(sejarah *models.Sejarah) error {
	return r.db.Create(sejarah).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *sejarahRepository) FindSejarah() (models.Sejarah, error) {
	var sejarah models.Sejarah
	err := r.db.First(&sejarah).Error
	return sejarah, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *sejarahRepository) UpdateSejarah(sejarah *models.Sejarah) error {
	return r.db.Save(sejarah).Error
}

func (r *sejarahRepository) DeleteSejarah() error {
	return r.db.Exec("DELETE FROM sejarah").Error
}