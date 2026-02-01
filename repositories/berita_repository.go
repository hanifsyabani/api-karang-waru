package repositories

import (
	"api-karang-waru/models"
	"strings"

	"gorm.io/gorm"
)

type BeritaRepository interface {
	CreateBerita(berita *models.Berita) error
	FindBerita(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.Berita, error)
	CountBeritaByCategory() (map[string]int64, error)
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
func (r *beritaRepository) FindBerita(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.Berita, error) {

	var berita []models.Berita
	offset := (page - 1) * limit
	query := r.db.Model(&models.Berita{})

	if search != "" {
		search = strings.ToLower(search)
		query = query.Where(
			"LOWER(title) LIKE ?",
			"%"+search+"%",
		)
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// sortBy → nama kolom
	// sortOrder → arah sorting
	query = query.Order(sortBy + " " + sortOrder)

	err := query.Limit(limit).Offset(offset).Find(&berita).Error
	return berita, err
}

func (r *beritaRepository) FindBeritaByID(id uint) (*models.Berita, error) {
	var berita models.Berita
	err := r.db.Find(&berita, id).Error
	return &berita, err
}

func (r *beritaRepository) CountBeritaByCategory() (map[string]int64, error) {
	type tempResult struct {
		Category string
		Total    int64
	}

	var dbresults []tempResult

	err := r.db.Model(&models.Berita{}).
		Select("category, COUNT(*) AS total").
		Group("category").
		Scan(&dbresults).Error

	if err != nil {
		return nil, err
	}
	kategoriList := []string{
		"umum",
		"kegiatan",
		"infrastruktur",
		"kesehatan",
		"pendidikan",
	}

	results := make(map[string]int64)
	for _, k := range kategoriList {
		results[k] = 0
	}
	var total int64

	// isi dari DB
	for _, r := range dbresults {
		results[r.Category] = r.Total
		total += r.Total

	}
	results["total"] = total

	return results, nil
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
