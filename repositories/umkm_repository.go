package repositories

import (
	"api-karang-waru/models"
	"strings"

	"gorm.io/gorm"
)

type UmkmRepository interface {
	CreateUmkm(umkm *models.Umkm) error
	FindUmkm(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
		status string,
	) ([]models.Umkm, error)
	FindUmkmByID(id uint) (*models.Umkm, error)
	CountStatus() (verified int64, unverified int64, err error)
	FindUmkmBySlug(slug string) (*models.Umkm, error)
	UpdateUmkm(umkm *models.Umkm) error
	DeleteUmkm(umkm *models.Umkm) error
}

type umkmRepository struct {
	db *gorm.DB
}

func NewUmkmRepository(db *gorm.DB) UmkmRepository {
	return &umkmRepository{db}
}

func (r *umkmRepository) CreateUmkm(umkm *models.Umkm) error {
	return r.db.Create(umkm).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *umkmRepository) FindUmkm(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
	status string,
) ([]models.Umkm, error) {
	var umkm []models.Umkm
	offset := (page - 1) * limit
	query := r.db.Model(&models.Umkm{})

	if search != "" {
		search = strings.ToLower(search)
		query = query.Where(
			"LOWER(nama_usaha) LIKE ?",
			"%"+search+"%",
		)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// sortBy → nama kolom
	// sortOrder → arah sorting
	query = query.Order(sortBy + " " + sortOrder)

	err := query.Limit(limit).Offset(offset).Find(&umkm).Error
	return umkm, err
}

func countByStatus(db *gorm.DB, status string) (int64, error) {
	var count int64

	err := db.Model(&models.Umkm{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (r *umkmRepository) CountStatus() (verified int64, unverified int64, err error) {
	verified, err = countByStatus(r.db, "verified")
	if err != nil {
		return
	}
	unverified, err = countByStatus(r.db, "unverified")
	if err != nil {
		return
	}

	return  
}
func (r *umkmRepository) FindUmkmByID(id uint) (*models.Umkm, error) {
	var umkm models.Umkm
	err := r.db.Find(&umkm, id).Error
	return &umkm, err
}

func (r *umkmRepository) FindUmkmBySlug(slug string) (*models.Umkm, error) {
	var umkm models.Umkm
	err := r.db.Where("slug = ?", slug).First(&umkm).Error
	return &umkm, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *umkmRepository) UpdateUmkm(umkm *models.Umkm) error {
	return r.db.Save(umkm).Error
}

func (r *umkmRepository) DeleteUmkm(umkm *models.Umkm) error {
	return r.db.Delete(umkm).Error

}
