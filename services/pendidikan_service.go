package services

import (
	"api-karang-waru/helpers"
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type PendidikanService interface {
	// Lembaga Pendidikan
	CreateLembagaPendidikan(req *requests.LembagaPendidikanRequest) (*models.LembagaPendidikan, error)
	GetAllLembagaPendidikan() ([]models.LembagaPendidikan, error)
	GetLembagaPendidikanByID(id uint) (*models.LembagaPendidikan, error)
	GetLembagaPendidikanBySlug(slug string) (*models.LembagaPendidikan, error)
	UpdateLembagaPendidikan(id uint, req *requests.LembagaPendidikanRequest) (*models.LembagaPendidikan, error)
	DeleteLembagaPendidikan(id uint) error

	// Statistik Pendidikan
	CreateStatistikPendidikan(req *requests.StatistikPendidikanRequest) (*models.StatistikPendidikan, error)
	GetAllStatistikPendidikan() ([]models.StatistikPendidikan, error)
	GetStatistikPendidikanByID(id uint) (*models.StatistikPendidikan, error)
	UpdateStatistikPendidikan(id uint, req *requests.StatistikPendidikanRequest) (*models.StatistikPendidikan, error)
	DeleteStatistikPendidikan(id uint) error

	// Program Pendidikan
	CreateProgramPendidikan(req *requests.ProgramPendidikanRequest) (*models.ProgramPendidikan, error)
	GetAllProgramPendidikan() ([]models.ProgramPendidikan, error)
	GetProgramPendidikanByID(id uint) (*models.ProgramPendidikan, error)
	UpdateProgramPendidikan(id uint, req *requests.ProgramPendidikanRequest) (*models.ProgramPendidikan, error)
	DeleteProgramPendidikan(id uint) error

	// Capaian Pendidikan
	CreateCapaianPendidikan(req *requests.CapaianPendidikanRequest) (*models.CapaianPendidikan, error)
	GetAllCapaianPendidikan() ([]models.CapaianPendidikan, error)
	GetCapaianPendidikanByID(id uint) (*models.CapaianPendidikan, error)
	UpdateCapaianPendidikan(id uint, req *requests.CapaianPendidikanRequest) (*models.CapaianPendidikan, error)
	DeleteCapaianPendidikan(id uint) error

	// Dokumentasi Pendidikan
	CreateDokumentasiPendidikan(req *requests.DokumentasiPendidikanRequest) (*models.DokumentasiPendidikan, error)
	GetAllDokumentasiPendidikan() ([]models.DokumentasiPendidikan, error)
	GetDokumentasiPendidikanByID(id uint) (*models.DokumentasiPendidikan, error)
	UpdateDokumentasiPendidikan(id uint, req *requests.DokumentasiPendidikanRequest) (*models.DokumentasiPendidikan, error)
	DeleteDokumentasiPendidikan(id uint) error
}

type pendidikanService struct {
	repo     repositories.PendidikanRepository
	validate *validator.Validate
}

func NewPendidikanService(repo repositories.PendidikanRepository) PendidikanService {
	return &pendidikanService{
		repo:     repo,
		validate: validator.New(),
	}
}

// ===============================
// 1. Lembaga Pendidikan
// ===============================
func (s *pendidikanService) CreateLembagaPendidikan(req *requests.LembagaPendidikanRequest) (*models.LembagaPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	lembaga := models.LembagaPendidikan{
		NamaSekolah:       req.NamaSekolah,
		Alamat:            req.Alamat,
		JenjangPendidikan: req.JenjangPendidikan,
		JumlahGuru:        req.JumlahGuru,
		JumlahSiswa:       req.JumlahSiswa,
		JumlahStaf:        req.JumlahGuru,
		Kontak:            req.Kontak,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
	}

	if err := s.repo.CreateLembagaPendidikan(&lembaga); err != nil {
		return nil, err
	}
	return &lembaga, nil
}

func (s *pendidikanService) GetAllLembagaPendidikan() ([]models.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikan()
}

func (s *pendidikanService) GetLembagaPendidikanByID(id uint) (*models.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikanByID(id)
}

func (s *pendidikanService) GetLembagaPendidikanBySlug(slug string) (*models.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikanBySlug(slug)
}

func (s *pendidikanService) UpdateLembagaPendidikan(id uint, req *requests.LembagaPendidikanRequest) (*models.LembagaPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	lembaga, err := s.repo.FindLembagaPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	lembaga.NamaSekolah = req.NamaSekolah
	lembaga.Alamat = req.Alamat
	lembaga.JenjangPendidikan = req.JenjangPendidikan
	lembaga.JumlahGuru = req.JumlahGuru
	lembaga.JumlahSiswa = req.JumlahSiswa
	lembaga.JumlahStaf = req.JumlahGuru
	lembaga.Kontak = req.Kontak
	lembaga.Latitude = req.Latitude
	lembaga.Longitude = req.Longitude

	if err := s.repo.UpdateLembagaPendidikan(lembaga); err != nil {
		return nil, err
	}
	return lembaga, nil
}

func (s *pendidikanService) DeleteLembagaPendidikan(id uint) error {
	lembaga, err := s.repo.FindLembagaPendidikanByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteLembagaPendidikan(lembaga)
}

// ===============================
// 2. Statistik Pendidikan
// ===============================
func (s *pendidikanService) CreateStatistikPendidikan(req *requests.StatistikPendidikanRequest) (*models.StatistikPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	stat := models.StatistikPendidikan{
		TidakSekolah:    req.TidakSekolah,
		Tahun:           req.Tahun,
		SD:              req.SD,
		SMP:             req.SMP,
		SMA:             req.SMA,
		PerguruanTinggi: req.PerguruanTinggi,
	}

	if err := s.repo.CreateStatistikPendidikan(&stat); err != nil {
		return nil, err
	}
	return &stat, nil
}

func (s *pendidikanService) GetAllStatistikPendidikan() ([]models.StatistikPendidikan, error) {
	return s.repo.FindStatistikPendidikan()
}

func (s *pendidikanService) GetStatistikPendidikanByID(id uint) (*models.StatistikPendidikan, error) {
	return s.repo.FindStatistikPendidikanByID(id)
}

func (s *pendidikanService) UpdateStatistikPendidikan(id uint, req *requests.StatistikPendidikanRequest) (*models.StatistikPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	stat, err := s.repo.FindStatistikPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	stat.Tahun = req.Tahun
	stat.TidakSekolah = req.TidakSekolah
	stat.PerguruanTinggi = req.PerguruanTinggi
	stat.SD = req.SD
	stat.SMA = req.SMA
	stat.SMP = req.SMP

	if err := s.repo.UpdateStatistikPendidikan(stat); err != nil {
		return nil, err
	}
	return stat, nil
}

func (s *pendidikanService) DeleteStatistikPendidikan(id uint) error {
	stat, err := s.repo.FindStatistikPendidikanByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteStatistikPendidikan(stat)
}

// ===============================
// 3. Program Pendidikan
// ===============================
func (s *pendidikanService) CreateProgramPendidikan(req *requests.ProgramPendidikanRequest) (*models.ProgramPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	parsedStartDate, err := helpers.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := helpers.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}

	prog := models.ProgramPendidikan{
		NamaProgram:    req.NamaProgram,
		Deskripsi:      req.Deskripsi,
		TanggalMulai:   parsedStartDate,
		TanggalSelesai: parsedEndDate,
	}

	if err := s.repo.CreateProgramPendidikan(&prog); err != nil {
		return nil, err
	}
	return &prog, nil
}

func (s *pendidikanService) GetAllProgramPendidikan() ([]models.ProgramPendidikan, error) {
	return s.repo.FindProgramPendidikan()
}

func (s *pendidikanService) GetProgramPendidikanByID(id uint) (*models.ProgramPendidikan, error) {
	return s.repo.FindProgramPendidikanByID(id)
}

func (s *pendidikanService) UpdateProgramPendidikan(id uint, req *requests.ProgramPendidikanRequest) (*models.ProgramPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	prog, err := s.repo.FindProgramPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	parsedStartDate, err := helpers.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := helpers.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}

	prog.NamaProgram = req.NamaProgram
	prog.Deskripsi = req.Deskripsi
	prog.TanggalMulai = parsedStartDate
	prog.TanggalSelesai = parsedEndDate
	prog.Status = req.Status

	if err := s.repo.UpdateProgramPendidikan(prog); err != nil {
		return nil, err
	}
	return prog, nil
}

func (s *pendidikanService) DeleteProgramPendidikan(id uint) error {
	prog, err := s.repo.FindProgramPendidikanByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteProgramPendidikan(prog)
}

// ===============================
// 4. Capaian Pendidikan
// ===============================
func (s *pendidikanService) CreateCapaianPendidikan(req *requests.CapaianPendidikanRequest) (*models.CapaianPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	cap := models.CapaianPendidikan{
		JumlahLulusan:      req.JumlahLulusan,
		TingkatPartisipasi: req.TingkatPartisipasi,
		TingkatAPK:         req.TingkatAPK,
		AngkaMelekHuruf:    req.AngkaMelekHuruf,
		Tahun:              req.Tahun,
	}

	if err := s.repo.CreateCapaianPendidikan(&cap); err != nil {
		return nil, err
	}
	return &cap, nil
}

func (s *pendidikanService) GetAllCapaianPendidikan() ([]models.CapaianPendidikan, error) {
	return s.repo.FindCapaianPendidikan()
}

func (s *pendidikanService) GetCapaianPendidikanByID(id uint) (*models.CapaianPendidikan, error) {
	return s.repo.FindCapaianPendidikanByID(id)
}

func (s *pendidikanService) UpdateCapaianPendidikan(id uint, req *requests.CapaianPendidikanRequest) (*models.CapaianPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	cap, err := s.repo.FindCapaianPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	cap.JumlahLulusan = req.JumlahLulusan
	cap.TingkatAPK = req.TingkatAPK
	cap.TingkatPartisipasi = req.TingkatPartisipasi
	cap.AngkaMelekHuruf = req.AngkaMelekHuruf
	cap.Tahun = req.Tahun

	if err := s.repo.UpdateCapaianPendidikan(cap); err != nil {
		return nil, err
	}
	return cap, nil
}

func (s *pendidikanService) DeleteCapaianPendidikan(id uint) error {
	cap, err := s.repo.FindCapaianPendidikanByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteCapaianPendidikan(cap)
}

// ===============================
// 5. Dokumentasi Pendidikan
// ===============================
func (s *pendidikanService) CreateDokumentasiPendidikan(req *requests.DokumentasiPendidikanRequest) (*models.DokumentasiPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	doc := models.DokumentasiPendidikan{
		Judul:        req.Judul,
		Keterangan:   req.Keterangan,
		FilePath:     req.FilePath,
		FotoKegiatan: req.FilePath,
		Tahun:        req.Tahun,
	}

	if err := s.repo.CreateDokumentasiPendidikan(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (s *pendidikanService) GetAllDokumentasiPendidikan() ([]models.DokumentasiPendidikan, error) {
	return s.repo.FindDokumentasiPendidikan()
}

func (s *pendidikanService) GetDokumentasiPendidikanByID(id uint) (*models.DokumentasiPendidikan, error) {
	return s.repo.FindDokumentasiPendidikanByID(id)
}

func (s *pendidikanService) UpdateDokumentasiPendidikan(id uint, req *requests.DokumentasiPendidikanRequest) (*models.DokumentasiPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	doc, err := s.repo.FindDokumentasiPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	doc.Judul = req.Judul
	doc.Keterangan = req.Keterangan
	doc.FilePath = req.FilePath
	doc.FotoKegiatan = req.FotoKegiatan
	doc.Tahun = req.Tahun

	if err := s.repo.UpdateDokumentasiPendidikan(doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *pendidikanService) DeleteDokumentasiPendidikan(id uint) error {
	doc, err := s.repo.FindDokumentasiPendidikanByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteDokumentasiPendidikan(doc)
}
