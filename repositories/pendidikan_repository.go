package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type PendidikanRepository interface {
	// Lembaga Pendidikan
	CreateLembagaPendidikan(lembaga *models.LembagaPendidikan) error
	FindLembagaPendidikan() ([]models.LembagaPendidikan, error)
	FindLembagaPendidikanByID(id uint) (*models.LembagaPendidikan, error)
	FindLembagaPendidikanBySlug(slug string) (*models.LembagaPendidikan, error)
	UpdateLembagaPendidikan(lembaga *models.LembagaPendidikan) error
	DeleteLembagaPendidikan(lembaga *models.LembagaPendidikan) error

	// Statistik Pendidikan
	CreateStatistikPendidikan(statistik *models.StatistikPendidikan) error
	FindStatistikPendidikan() ([]models.StatistikPendidikan, error)
	FindStatistikPendidikanByID(id uint) (*models.StatistikPendidikan, error)
	UpdateStatistikPendidikan(statistik *models.StatistikPendidikan) error
	DeleteStatistikPendidikan(statistik *models.StatistikPendidikan) error

	// Program Pendidikan
	CreateProgramPendidikan(program *models.ProgramPendidikan) error
	FindProgramPendidikan() ([]models.ProgramPendidikan, error)
	FindProgramPendidikanByID(id uint) (*models.ProgramPendidikan, error)
	UpdateProgramPendidikan(program *models.ProgramPendidikan) error
	DeleteProgramPendidikan(program *models.ProgramPendidikan) error

	// Capaian Pendidikan
	CreateCapaianPendidikan(capaian *models.CapaianPendidikan) error
	FindCapaianPendidikan() ([]models.CapaianPendidikan, error)
	FindCapaianPendidikanByID(id uint) (*models.CapaianPendidikan, error)
	UpdateCapaianPendidikan(capaian *models.CapaianPendidikan) error
	DeleteCapaianPendidikan(capaian *models.CapaianPendidikan) error

	// Dokumentasi Pendidikan
	CreateDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error
	FindDokumentasiPendidikan() ([]models.DokumentasiPendidikan, error)
	FindDokumentasiPendidikanByID(id uint) (*models.DokumentasiPendidikan, error)
	UpdateDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error
	DeleteDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error
}

type pendidikanRepository struct {
	db *gorm.DB
}

func NewPendidikanRepository(db *gorm.DB) PendidikanRepository {
	return &pendidikanRepository{db}
}

//
// ===============================
// 1. Lembaga Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateLembagaPendidikan(lembaga *models.LembagaPendidikan) error {
	return r.db.Create(lembaga).Error
}

func (r *pendidikanRepository) FindLembagaPendidikan() ([]models.LembagaPendidikan, error) {
	var lembaga []models.LembagaPendidikan
	err := r.db.Find(&lembaga).Error
	return lembaga, err
}

func (r *pendidikanRepository) FindLembagaPendidikanByID(id uint) (*models.LembagaPendidikan, error) {
	var lembaga models.LembagaPendidikan
	err := r.db.First(&lembaga, id).Error
	return &lembaga, err
}

func (r *pendidikanRepository) FindLembagaPendidikanBySlug(slug string) (*models.LembagaPendidikan, error) {
	var lembaga models.LembagaPendidikan
	err := r.db.Where("slug = ?", slug).First(&lembaga).Error
	return &lembaga, err
}

func (r *pendidikanRepository) UpdateLembagaPendidikan(lembaga *models.LembagaPendidikan) error {
	return r.db.Save(lembaga).Error
}

func (r *pendidikanRepository) DeleteLembagaPendidikan(lembaga *models.LembagaPendidikan) error {
	return r.db.Delete(lembaga).Error
}

//
// ===============================
// 2. Statistik Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateStatistikPendidikan(statistik *models.StatistikPendidikan) error {
	return r.db.Create(statistik).Error
}

func (r *pendidikanRepository) FindStatistikPendidikan() ([]models.StatistikPendidikan, error) {
	var statistik []models.StatistikPendidikan
	err := r.db.Find(&statistik).Error
	return statistik, err
}

func (r *pendidikanRepository) FindStatistikPendidikanByID(id uint) (*models.StatistikPendidikan, error) {
	var statistik models.StatistikPendidikan
	err := r.db.First(&statistik, id).Error
	return &statistik, err
}

func (r *pendidikanRepository) UpdateStatistikPendidikan(statistik *models.StatistikPendidikan) error {
	return r.db.Save(statistik).Error
}

func (r *pendidikanRepository) DeleteStatistikPendidikan(statistik *models.StatistikPendidikan) error {
	return r.db.Delete(statistik).Error
}

//
// ===============================
// 3. Program Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateProgramPendidikan(program *models.ProgramPendidikan) error {
	return r.db.Create(program).Error
}

func (r *pendidikanRepository) FindProgramPendidikan() ([]models.ProgramPendidikan, error) {
	var program []models.ProgramPendidikan
	err := r.db.Find(&program).Error
	return program, err
}

func (r *pendidikanRepository) FindProgramPendidikanByID(id uint) (*models.ProgramPendidikan, error) {
	var program models.ProgramPendidikan
	err := r.db.First(&program, id).Error
	return &program, err
}

func (r *pendidikanRepository) UpdateProgramPendidikan(program *models.ProgramPendidikan) error {
	return r.db.Save(program).Error
}

func (r *pendidikanRepository) DeleteProgramPendidikan(program *models.ProgramPendidikan) error {
	return r.db.Delete(program).Error
}

//
// ===============================
// 4. Capaian Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateCapaianPendidikan(capaian *models.CapaianPendidikan) error {
	return r.db.Create(capaian).Error
}

func (r *pendidikanRepository) FindCapaianPendidikan() ([]models.CapaianPendidikan, error) {
	var capaian []models.CapaianPendidikan
	err := r.db.Find(&capaian).Error
	return capaian, err
}

func (r *pendidikanRepository) FindCapaianPendidikanByID(id uint) (*models.CapaianPendidikan, error) {
	var capaian models.CapaianPendidikan
	err := r.db.First(&capaian, id).Error
	return &capaian, err
}

func (r *pendidikanRepository) UpdateCapaianPendidikan(capaian *models.CapaianPendidikan) error {
	return r.db.Save(capaian).Error
}

func (r *pendidikanRepository) DeleteCapaianPendidikan(capaian *models.CapaianPendidikan) error {
	return r.db.Delete(capaian).Error
}

//
// ===============================
// 5. Dokumentasi & Laporan Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error {
	return r.db.Create(dok).Error
}

func (r *pendidikanRepository) FindDokumentasiPendidikan() ([]models.DokumentasiPendidikan, error) {
	var dok []models.DokumentasiPendidikan
	err := r.db.Find(&dok).Error
	return dok, err
}

func (r *pendidikanRepository) FindDokumentasiPendidikanByID(id uint) (*models.DokumentasiPendidikan, error) {
	var dok models.DokumentasiPendidikan
	err := r.db.First(&dok, id).Error
	return &dok, err
}

func (r *pendidikanRepository) UpdateDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error {
	return r.db.Save(dok).Error
}

func (r *pendidikanRepository) DeleteDokumentasiPendidikan(dok *models.DokumentasiPendidikan) error {
	return r.db.Delete(dok).Error
}
