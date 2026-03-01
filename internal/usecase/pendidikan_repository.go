package usecase

import (
	"api-karang-waru/internal/domain"

	"gorm.io/gorm"
)

type PendidikanRepository interface {
	// Lembaga Pendidikan
	CreateLembagaPendidikan(lembaga *domain.LembagaPendidikan) error
	FindLembagaPendidikan() ([]domain.LembagaPendidikan, error)
	FindLembagaPendidikanByID(id uint) (*domain.LembagaPendidikan, error)
	FindLembagaPendidikanBySlug(slug string) (*domain.LembagaPendidikan, error)
	UpdateLembagaPendidikan(lembaga *domain.LembagaPendidikan) error
	DeleteLembagaPendidikan(lembaga *domain.LembagaPendidikan) error

	// Statistik Pendidikan
	CreateStatistikPendidikan(statistik *domain.StatistikPendidikan) error
	FindStatistikPendidikan() ([]domain.StatistikPendidikan, error)
	FindStatistikPendidikanByID(id uint) (*domain.StatistikPendidikan, error)
	UpdateStatistikPendidikan(statistik *domain.StatistikPendidikan) error
	DeleteStatistikPendidikan(statistik *domain.StatistikPendidikan) error

	// Program Pendidikan
	CreateProgramPendidikan(program *domain.ProgramPendidikan) error
	FindProgramPendidikan() ([]domain.ProgramPendidikan, error)
	FindProgramPendidikanByID(id uint) (*domain.ProgramPendidikan, error)
	UpdateProgramPendidikan(program *domain.ProgramPendidikan) error
	DeleteProgramPendidikan(program *domain.ProgramPendidikan) error

	// Capaian Pendidikan
	CreateCapaianPendidikan(capaian *domain.CapaianPendidikan) error
	FindCapaianPendidikan() ([]domain.CapaianPendidikan, error)
	FindCapaianPendidikanByID(id uint) (*domain.CapaianPendidikan, error)
	UpdateCapaianPendidikan(capaian *domain.CapaianPendidikan) error
	DeleteCapaianPendidikan(capaian *domain.CapaianPendidikan) error

	// Dokumentasi Pendidikan
	CreateDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error
	FindDokumentasiPendidikan() ([]domain.DokumentasiPendidikan, error)
	FindDokumentasiPendidikanByID(id uint) (*domain.DokumentasiPendidikan, error)
	UpdateDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error
	DeleteDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error
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
func (r *pendidikanRepository) CreateLembagaPendidikan(lembaga *domain.LembagaPendidikan) error {
	return r.db.Create(lembaga).Error
}

func (r *pendidikanRepository) FindLembagaPendidikan() ([]domain.LembagaPendidikan, error) {
	var lembaga []domain.LembagaPendidikan
	err := r.db.Find(&lembaga).Error
	return lembaga, err
}

func (r *pendidikanRepository) FindLembagaPendidikanByID(id uint) (*domain.LembagaPendidikan, error) {
	var lembaga domain.LembagaPendidikan
	err := r.db.First(&lembaga, id).Error
	return &lembaga, err
}

func (r *pendidikanRepository) FindLembagaPendidikanBySlug(slug string) (*domain.LembagaPendidikan, error) {
	var lembaga domain.LembagaPendidikan
	err := r.db.Where("slug = ?", slug).First(&lembaga).Error
	return &lembaga, err
}

func (r *pendidikanRepository) UpdateLembagaPendidikan(lembaga *domain.LembagaPendidikan) error {
	return r.db.Save(lembaga).Error
}

func (r *pendidikanRepository) DeleteLembagaPendidikan(lembaga *domain.LembagaPendidikan) error {
	return r.db.Delete(lembaga).Error
}

//
// ===============================
// 2. Statistik Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateStatistikPendidikan(statistik *domain.StatistikPendidikan) error {
	return r.db.Create(statistik).Error
}

func (r *pendidikanRepository) FindStatistikPendidikan() ([]domain.StatistikPendidikan, error) {
	var statistik []domain.StatistikPendidikan
	err := r.db.Find(&statistik).Error
	return statistik, err
}

func (r *pendidikanRepository) FindStatistikPendidikanByID(id uint) (*domain.StatistikPendidikan, error) {
	var statistik domain.StatistikPendidikan
	err := r.db.First(&statistik, id).Error
	return &statistik, err
}

func (r *pendidikanRepository) UpdateStatistikPendidikan(statistik *domain.StatistikPendidikan) error {
	return r.db.Save(statistik).Error
}

func (r *pendidikanRepository) DeleteStatistikPendidikan(statistik *domain.StatistikPendidikan) error {
	return r.db.Delete(statistik).Error
}

//
// ===============================
// 3. Program Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateProgramPendidikan(program *domain.ProgramPendidikan) error {
	return r.db.Create(program).Error
}

func (r *pendidikanRepository) FindProgramPendidikan() ([]domain.ProgramPendidikan, error) {
	var program []domain.ProgramPendidikan
	err := r.db.Find(&program).Error
	return program, err
}

func (r *pendidikanRepository) FindProgramPendidikanByID(id uint) (*domain.ProgramPendidikan, error) {
	var program domain.ProgramPendidikan
	err := r.db.First(&program, id).Error
	return &program, err
}

func (r *pendidikanRepository) UpdateProgramPendidikan(program *domain.ProgramPendidikan) error {
	return r.db.Save(program).Error
}

func (r *pendidikanRepository) DeleteProgramPendidikan(program *domain.ProgramPendidikan) error {
	return r.db.Delete(program).Error
}

//
// ===============================
// 4. Capaian Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateCapaianPendidikan(capaian *domain.CapaianPendidikan) error {
	return r.db.Create(capaian).Error
}

func (r *pendidikanRepository) FindCapaianPendidikan() ([]domain.CapaianPendidikan, error) {
	var capaian []domain.CapaianPendidikan
	err := r.db.Find(&capaian).Error
	return capaian, err
}

func (r *pendidikanRepository) FindCapaianPendidikanByID(id uint) (*domain.CapaianPendidikan, error) {
	var capaian domain.CapaianPendidikan
	err := r.db.First(&capaian, id).Error
	return &capaian, err
}

func (r *pendidikanRepository) UpdateCapaianPendidikan(capaian *domain.CapaianPendidikan) error {
	return r.db.Save(capaian).Error
}

func (r *pendidikanRepository) DeleteCapaianPendidikan(capaian *domain.CapaianPendidikan) error {
	return r.db.Delete(capaian).Error
}

//
// ===============================
// 5. Dokumentasi & Laporan Pendidikan
// ===============================
//
func (r *pendidikanRepository) CreateDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error {
	return r.db.Create(dok).Error
}

func (r *pendidikanRepository) FindDokumentasiPendidikan() ([]domain.DokumentasiPendidikan, error) {
	var dok []domain.DokumentasiPendidikan
	err := r.db.Find(&dok).Error
	return dok, err
}

func (r *pendidikanRepository) FindDokumentasiPendidikanByID(id uint) (*domain.DokumentasiPendidikan, error) {
	var dok domain.DokumentasiPendidikan
	err := r.db.First(&dok, id).Error
	return &dok, err
}

func (r *pendidikanRepository) UpdateDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error {
	return r.db.Save(dok).Error
}

func (r *pendidikanRepository) DeleteDokumentasiPendidikan(dok *domain.DokumentasiPendidikan) error {
	return r.db.Delete(dok).Error
}

