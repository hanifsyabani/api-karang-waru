package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type VisiMisiRepository interface {
	CreateVisiMisi(visiMisi *domain.VisiMisi) error
	FindVisiMisi() (domain.VisiMisi, error)
	UpdateVisiMisi(visiMisi *domain.VisiMisi) error
	DeleteVisiMisi() error
}

type visiMisiRepository struct {
	db *gorm.DB
}

func NewVisiMisiRepository(db *gorm.DB) VisiMisiRepository {
	return &visiMisiRepository{db}
}

func (r *visiMisiRepository) CreateVisiMisi(visiMisi *domain.VisiMisi) error {
	return r.db.Create(visiMisi).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *visiMisiRepository) FindVisiMisi() (domain.VisiMisi, error) {
	var visiMisi domain.VisiMisi
	err := r.db.First(&visiMisi).Error
	return visiMisi, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *visiMisiRepository) UpdateVisiMisi(visiMisi *domain.VisiMisi) error {
	return r.db.Save(visiMisi).Error
}

func (r *visiMisiRepository) DeleteVisiMisi() error {
	return r.db.Exec("DELETE FROM visi_misi").Error
}
