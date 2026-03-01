package usecase

import (
	"api-karang-waru/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type BeritaRepository interface {
	CreateBerita(berita *domain.Berita) error
	FindBerita(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.Berita, error)
	CountBeritaByCategory() (map[string]int64, error)
	FindBeritaByID(id uint) (*domain.Berita, error)
	FindBeritaBySlug(slug string) (*domain.Berita, error)
	UpdateBerita(berita *domain.Berita) error
	DeleteBerita(berita *domain.Berita) error
}

type beritaRepository struct {
	db *gorm.DB
}

func NewBeritaRepository(db *gorm.DB) BeritaRepository {
	return &beritaRepository{db}
}

func (r *beritaRepository) CreateBerita(berita *domain.Berita) error {
	return r.db.Create(berita).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *beritaRepository) FindBerita(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.Berita, error) {

	var berita []domain.Berita
	offset := (page - 1) * limit
	query := r.db.Model(&domain.Berita{})

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

func (r *beritaRepository) FindBeritaByID(id uint) (*domain.Berita, error) {
	var berita domain.Berita
	err := r.db.Find(&berita, id).Error
	return &berita, err
}

func (r *beritaRepository) CountBeritaByCategory() (map[string]int64, error) {
	type tempResult struct {
		Category string
		Total    int64
	}

	var dbresults []tempResult

	err := r.db.Model(&domain.Berita{}).
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

func (r *beritaRepository) FindBeritaBySlug(slug string) (*domain.Berita, error) {
	var berita domain.Berita
	err := r.db.Where("slug = ?", slug).First(&berita).Error
	return &berita, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *beritaRepository) UpdateBerita(berita *domain.Berita) error {
	return r.db.Save(berita).Error
}

func (r *beritaRepository) DeleteBerita(berita *domain.Berita) error {
	return r.db.Delete(berita).Error

}

