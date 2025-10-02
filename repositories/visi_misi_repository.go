package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type VisiMisiRepository interface {
	CreateVisiMisi(visiMisi *models.VisiMisi) error
	FindVisiMisi() (models.VisiMisi, error)
	UpdateVisiMisi(visiMisi *models.VisiMisi) error
	DeleteVisiMisi() error
}

type visiMisiRepository struct {
	db *gorm.DB
}

func NewVisiMisiRepository(db *gorm.DB) VisiMisiRepository {
	return &visiMisiRepository{db}
}

func (r *visiMisiRepository) CreateVisiMisi(visiMisi *models.VisiMisi) error {
	return r.db.Create(visiMisi).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *visiMisiRepository) FindVisiMisi() (models.VisiMisi, error) {
	var visiMisi models.VisiMisi
	err := r.db.First(&visiMisi).Error
	return visiMisi, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *visiMisiRepository) UpdateVisiMisi(visiMisi *models.VisiMisi) error {
	return r.db.Save(visiMisi).Error
}

func (r *visiMisiRepository) DeleteVisiMisi() error {
	return r.db.Exec("DELETE FROM visi_misi").Error
}