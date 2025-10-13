package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type BeritaRepository interface {
	CreateBerita(berita *models.Berita) error
	FindBerita() ([]models.Berita, error)
	FindBeritaByID(id uint) (*models.Berita, error)
	FindBeritaBySlug(slug string) (*models.Berita, error)
	UpdateBerita(berita *models.Berita) error
	DeleteBerita(berita *models.Berita) error
}

type beritaRepository struct {
	db *gorm.DB
}

func NewBeritaRepository(db *gorm.DB) BeritaRepository {
	return &beritaRepository{db}
}

func (r *beritaRepository) CreateBerita(berita *models.Berita) error {
	return r.db.Create(berita).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *beritaRepository) FindBerita() ([]models.Berita, error) {
	var berita []models.Berita
	err := r.db.Find(&berita).Error
	return berita, err
}

func (r *beritaRepository) FindBeritaByID(id uint) (*models.Berita, error) {
	var berita models.Berita
	err := r.db.Find(&berita, id).Error
	return &berita, err
}

func (r *beritaRepository) FindBeritaBySlug(slug string) (*models.Berita, error) {
	var berita models.Berita
	err := r.db.Where("slug = ?", slug).First(&berita).Error
	return &berita, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *beritaRepository) UpdateBerita(berita *models.Berita) error {
	return r.db.Save(berita).Error
}

func (r *beritaRepository) DeleteBerita(berita *models.Berita) error {
	return r.db.Delete(berita).Error

}
