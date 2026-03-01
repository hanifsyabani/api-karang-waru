package usecase

import (
	"api-karang-waru/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type PendudukRepository interface {
	CreatePenduduk(penduduk *domain.Penduduk) error
	CountPenduduk() (total int64, lakiLaki int64, perempuan int64, kartuKeluarga int64, err error)
	FindPenduduk(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.Penduduk, error)
	FindPendudukByID(id uint) (*domain.Penduduk, error)
	IsNIKExists(nik string) (bool, error)
	UpdatePenduduk(penduduk *domain.Penduduk) error
	DeletePenduduk(penduduk *domain.Penduduk) error
}

type pendudukRepository struct {
	db *gorm.DB
}

func NewPendudukRepository(db *gorm.DB) PendudukRepository {
	return &pendudukRepository{db}
}

func (r *pendudukRepository) CreatePenduduk(penduduk *domain.Penduduk) error {
	return r.db.Create(penduduk).Error
}

// []domain.Demografis artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.Demografis.
func (r *pendudukRepository) FindPenduduk(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.Penduduk, error) {
	var penduduk []domain.Penduduk

	offset := (page - 1) * limit
	query := r.db.Model(&domain.Penduduk{})

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

	err:= db.Model(&domain.Penduduk{}).Where("jenis_kelamin = ?", gender).Count(&count).Error
	return count, err
}

func (r *pendudukRepository) CountPenduduk() (total int64, lakiLaki int64, perempuan int64, kartuKeluarga int64, err error) {
	if err = r.db.Model(&domain.Penduduk{}).Count(&total).Error; err != nil {
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

	if err = r.db.Model(&domain.Penduduk{}).Distinct("no_kk").Count(&kartuKeluarga).Error; err != nil {
		return
	}

	return
}

func (r *pendudukRepository) FindPendudukByID(id uint) (*domain.Penduduk, error) {
	var penduduk domain.Penduduk
	err := r.db.Find(&penduduk, id).Error
	return &penduduk, err
}

func (r *pendudukRepository) IsNIKExists(nik string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Penduduk{}).
		Where("nik = ? AND deleted_at IS NULL", nik).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.ProfilDesa
func (r *pendudukRepository) UpdatePenduduk(penduduk *domain.Penduduk) error {
	return r.db.Save(penduduk).Error
}

func (r *pendudukRepository) DeletePenduduk(penduduk *domain.Penduduk) error {
	return r.db.Delete(penduduk).Error

}

