package usecase

import (
	"api-karang-waru/pkg/utils"
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type PendidikanService interface {
	// Lembaga Pendidikan
	CreateLembagaPendidikan(req *types.LembagaPendidikanRequest) (*domain.LembagaPendidikan, error)
	GetAllLembagaPendidikan() ([]domain.LembagaPendidikan, error)
	GetLembagaPendidikanByID(id uint) (*domain.LembagaPendidikan, error)
	GetLembagaPendidikanBySlug(slug string) (*domain.LembagaPendidikan, error)
	UpdateLembagaPendidikan(id uint, req *types.LembagaPendidikanRequest) (*domain.LembagaPendidikan, error)
	DeleteLembagaPendidikan(id uint) error

	// Statistik Pendidikan
	CreateStatistikPendidikan(req *types.StatistikPendidikanRequest) (*domain.StatistikPendidikan, error)
	GetAllStatistikPendidikan() ([]domain.StatistikPendidikan, error)
	GetStatistikPendidikanByID(id uint) (*domain.StatistikPendidikan, error)
	UpdateStatistikPendidikan(id uint, req *types.StatistikPendidikanRequest) (*domain.StatistikPendidikan, error)
	DeleteStatistikPendidikan(id uint) error

	// Program Pendidikan
	CreateProgramPendidikan(req *types.ProgramPendidikanRequest) (*domain.ProgramPendidikan, error)
	GetAllProgramPendidikan() ([]domain.ProgramPendidikan, error)
	GetProgramPendidikanByID(id uint) (*domain.ProgramPendidikan, error)
	UpdateProgramPendidikan(id uint, req *types.ProgramPendidikanRequest) (*domain.ProgramPendidikan, error)
	DeleteProgramPendidikan(id uint) error

	// Capaian Pendidikan
	CreateCapaianPendidikan(req *types.CapaianPendidikanRequest) (*domain.CapaianPendidikan, error)
	GetAllCapaianPendidikan() ([]domain.CapaianPendidikan, error)
	GetCapaianPendidikanByID(id uint) (*domain.CapaianPendidikan, error)
	UpdateCapaianPendidikan(id uint, req *types.CapaianPendidikanRequest) (*domain.CapaianPendidikan, error)
	DeleteCapaianPendidikan(id uint) error

	// Dokumentasi Pendidikan
	CreateDokumentasiPendidikan(req *types.DokumentasiPendidikanRequest) (*domain.DokumentasiPendidikan, error)
	GetAllDokumentasiPendidikan() ([]domain.DokumentasiPendidikan, error)
	GetDokumentasiPendidikanByID(id uint) (*domain.DokumentasiPendidikan, error)
	UpdateDokumentasiPendidikan(id uint, req *types.DokumentasiPendidikanRequest) (*domain.DokumentasiPendidikan, error)
	DeleteDokumentasiPendidikan(id uint) error
}

type pendidikanService struct {
	repo     PendidikanRepository
	validate *validator.Validate
}

func NewPendidikanService(repo PendidikanRepository) PendidikanService {
	return &pendidikanService{
		repo:     repo,
		validate: validator.New(),
	}
}

// ===============================
// 1. Lembaga Pendidikan
// ===============================
func (s *pendidikanService) CreateLembagaPendidikan(req *types.LembagaPendidikanRequest) (*domain.LembagaPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	lembaga := domain.LembagaPendidikan{
		NamaSekolah:       req.NamaSekolah,
		Alamat:            req.Alamat,
		JenjangPendidikan: req.JenjangPendidikan,
		JumlahGuru:        req.JumlahGuru,
		JumlahSiswa:       req.JumlahSiswa,
		JumlahStaf:        req.JumlahGuru,
		Kontak:            req.Kontak,
	}

	if err := s.repo.CreateLembagaPendidikan(&lembaga); err != nil {
		return nil, err
	}
	return &lembaga, nil
}

func (s *pendidikanService) GetAllLembagaPendidikan() ([]domain.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikan()
}

func (s *pendidikanService) GetLembagaPendidikanByID(id uint) (*domain.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikanByID(id)
}

func (s *pendidikanService) GetLembagaPendidikanBySlug(slug string) (*domain.LembagaPendidikan, error) {
	return s.repo.FindLembagaPendidikanBySlug(slug)
}

func (s *pendidikanService) UpdateLembagaPendidikan(id uint, req *types.LembagaPendidikanRequest) (*domain.LembagaPendidikan, error) {
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
func (s *pendidikanService) CreateStatistikPendidikan(req *types.StatistikPendidikanRequest) (*domain.StatistikPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	stat := domain.StatistikPendidikan{
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

func (s *pendidikanService) GetAllStatistikPendidikan() ([]domain.StatistikPendidikan, error) {
	return s.repo.FindStatistikPendidikan()
}

func (s *pendidikanService) GetStatistikPendidikanByID(id uint) (*domain.StatistikPendidikan, error) {
	return s.repo.FindStatistikPendidikanByID(id)
}

func (s *pendidikanService) UpdateStatistikPendidikan(id uint, req *types.StatistikPendidikanRequest) (*domain.StatistikPendidikan, error) {
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
func (s *pendidikanService) CreateProgramPendidikan(req *types.ProgramPendidikanRequest) (*domain.ProgramPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	parsedStartDate, err := utils.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := utils.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}

	prog := domain.ProgramPendidikan{
		NamaProgram:    req.NamaProgram,
		Deskripsi:      req.Deskripsi,
		TanggalMulai:   parsedStartDate,
		TanggalSelesai: parsedEndDate,
		Status:         req.Status,
	}

	if err := s.repo.CreateProgramPendidikan(&prog); err != nil {
		return nil, err
	}
	return &prog, nil
}

func (s *pendidikanService) GetAllProgramPendidikan() ([]domain.ProgramPendidikan, error) {
	return s.repo.FindProgramPendidikan()
}

func (s *pendidikanService) GetProgramPendidikanByID(id uint) (*domain.ProgramPendidikan, error) {
	return s.repo.FindProgramPendidikanByID(id)
}

func (s *pendidikanService) UpdateProgramPendidikan(id uint, req *types.ProgramPendidikanRequest) (*domain.ProgramPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	prog, err := s.repo.FindProgramPendidikanByID(id)
	if err != nil {
		return nil, err
	}

	parsedStartDate, err := utils.ParseTimeHuman(req.TanggalMulai)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := utils.ParseTimeHuman(req.TanggalMulai)
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
func (s *pendidikanService) CreateCapaianPendidikan(req *types.CapaianPendidikanRequest) (*domain.CapaianPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	cap := domain.CapaianPendidikan{
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

func (s *pendidikanService) GetAllCapaianPendidikan() ([]domain.CapaianPendidikan, error) {
	return s.repo.FindCapaianPendidikan()
}

func (s *pendidikanService) GetCapaianPendidikanByID(id uint) (*domain.CapaianPendidikan, error) {
	return s.repo.FindCapaianPendidikanByID(id)
}

func (s *pendidikanService) UpdateCapaianPendidikan(id uint, req *types.CapaianPendidikanRequest) (*domain.CapaianPendidikan, error) {
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
func (s *pendidikanService) CreateDokumentasiPendidikan(req *types.DokumentasiPendidikanRequest) (*domain.DokumentasiPendidikan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	doc := domain.DokumentasiPendidikan{
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

func (s *pendidikanService) GetAllDokumentasiPendidikan() ([]domain.DokumentasiPendidikan, error) {
	return s.repo.FindDokumentasiPendidikan()
}

func (s *pendidikanService) GetDokumentasiPendidikanByID(id uint) (*domain.DokumentasiPendidikan, error) {
	return s.repo.FindDokumentasiPendidikanByID(id)
}

func (s *pendidikanService) UpdateDokumentasiPendidikan(id uint, req *types.DokumentasiPendidikanRequest) (*domain.DokumentasiPendidikan, error) {
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

