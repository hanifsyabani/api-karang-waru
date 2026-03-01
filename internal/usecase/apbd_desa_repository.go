package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type ApbdDesaRepository interface {
	CreateApbd(apbd *domain.APBDDesa) error
	FindApbd() ([]domain.APBDDesa, error)
	FindApbdByID(id uint) (*domain.APBDDesa, error)
	UpdateApbd(apbd *domain.APBDDesa) error
	DeleteApbd(apbd *domain.APBDDesa) error
}

type apbdRepository struct {
	db *gorm.DB
}

func NewApbdRepository(db *gorm.DB) ApbdDesaRepository {
	return &apbdRepository{db}
}

func (r *apbdRepository) CreateApbd(apbd *domain.APBDDesa) error {
	return r.db.Create(apbd).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *apbdRepository) FindApbd() ([]domain.APBDDesa, error) {
	var apbd []domain.APBDDesa
	err := r.db.Find(&apbd).Error
	return apbd, err
}

func (r *apbdRepository) FindApbdByID(id uint) (*domain.APBDDesa, error) {
	var apbd domain.APBDDesa
	err := r.db.Find(&apbd, id).Error
	return &apbd, err
}


// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *apbdRepository) UpdateApbd(apbd *domain.APBDDesa) error {
	return r.db.Save(apbd).Error
}

func (r *apbdRepository) DeleteApbd(apbd *domain.APBDDesa) error {
	return r.db.Delete(apbd).Error

}

