package repositories

import (
	"api-karang-waru/models"
	"strings"

	"gorm.io/gorm"
)

type PendudukRepository interface {
	CreatePenduduk(penduduk *models.Penduduk) error
	CountPenduduk() (total int64, lakiLaki int64, perempuan int64, kartuKeluarga int64, err error)
	FindPenduduk(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.Penduduk, error)
	FindPendudukByID(id uint) (*models.Penduduk, error)
	IsNIKExists(nik string) (bool, error)
	UpdatePenduduk(penduduk *models.Penduduk) error
	DeletePenduduk(penduduk *models.Penduduk) error
}

type pendudukRepository struct {
	db *gorm.DB
}

func NewPendudukRepository(db *gorm.DB) PendudukRepository {
	return &pendudukRepository{db}
}

func (r *pendudukRepository) CreatePenduduk(penduduk *models.Penduduk) error {
	return r.db.Create(penduduk).Error
}

// []models.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.Demografis.
func (r *pendudukRepository) FindPenduduk(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.Penduduk, error) {
	var penduduk []models.Penduduk

	offset := (page - 1) * limit
	query := r.db.Model(&models.Penduduk{})

	if search != "" {
		search = strings.ToLower(search)
		query = query.Where(
			"LOWER(nik) LIKE ? OR LOWER(nama_lengkap) LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	query = query.Order(sortBy + " " + sortOrder)

	err := query.Limit(limit).Offset(offset).Find(&penduduk).Error
	return penduduk, err
}

func countByGender(db *gorm.DB, gender string) (int64, error) {
	var count int64

	err:= db.Model(&models.Penduduk{}).Where("jenis_kelamin = ?", gender).Count(&count).Error
	return count, err
}

func (r *pendudukRepository) CountPenduduk() (total int64, lakiLaki int64, perempuan int64, kartuKeluarga int64, err error) {
	if err = r.db.Model(&models.Penduduk{}).Count(&total).Error; err != nil {
		return
	}

	lakiLaki, err = countByGender(r.db, "Laki-laki")
	if err != nil {
		return
	}

	perempuan, err = countByGender(r.db, "Perempuan")
	if err != nil {
		return
	}

	if err = r.db.Model(&models.Penduduk{}).Distinct("no_kk").Count(&kartuKeluarga).Error; err != nil {
		return
	}

	return
}

func (r *pendudukRepository) FindPendudukByID(id uint) (*models.Penduduk, error) {
	var penduduk models.Penduduk
	err := r.db.Find(&penduduk, id).Error
	return &penduduk, err
}

func (r *pendudukRepository) IsNIKExists(nik string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Penduduk{}).
		Where("nik = ? AND deleted_at IS NULL", nik).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.ProfilDesa
func (r *pendudukRepository) UpdatePenduduk(penduduk *models.Penduduk) error {
	return r.db.Save(penduduk).Error
}

func (r *pendudukRepository) DeletePenduduk(penduduk *models.Penduduk) error {
	return r.db.Delete(penduduk).Error

}
