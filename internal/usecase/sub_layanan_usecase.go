package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

// ==================== SUB LAYANAN SERVICE ====================

type SubLayananService interface {
	CreateSubLayanan(req *types.SubLayananRequest) (*domain.SubLayananDesa, error)
	GetAllSubLayanan(
		layananDesaID uint,
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.SubLayananDesa, error)
	GetSubLayananByID(id uint) (*domain.SubLayananDesa, error)
	GetSubLayananByLayananID(layananDesaID uint) ([]domain.SubLayananDesa, error)
	UpdateSubLayanan(id uint, req *types.SubLayananRequest) (*domain.SubLayananDesa, error)
	DeleteSubLayanan(id uint) error
}

type subLayananService struct {
	repository        SubLayananRepository
	layananRepository LayananRepository
	validate          *validator.Validate
}

func NewSubLayananService(
	repository SubLayananRepository,
	layananRepository LayananRepository,
) SubLayananService {
	return &subLayananService{
		repository:        repository,
		layananRepository: layananRepository,
		validate:          validator.New(),
	}
}

func (s *subLayananService) CreateSubLayanan(req *types.SubLayananRequest) (*domain.SubLayananDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Validasi layanan desa exists
	_, err := s.layananRepository.FindLayananByID(req.LayananDesaID)
	if err != nil {
		return nil, errors.New("layanan desa tidak ditemukan")
	}

	subLayanan := domain.SubLayananDesa{
		LayananDesaID: req.LayananDesaID,
		Nama:          req.Nama,
		Persyaratan:   req.Persyaratan,
		Template:      req.Template,
		Aktif:         req.Aktif,
	}

	err = s.repository.CreateSubLayanan(&subLayanan)
	return &subLayanan, err
}

func (s *subLayananService) GetAllSubLayanan(
	layananDesaID uint,
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.SubLayananDesa, error) {
	return s.repository.FindSubLayanan(layananDesaID, search, page, limit, sortBy, sortOrder)
}

func (s *subLayananService) GetSubLayananByID(id uint) (*domain.SubLayananDesa, error) {
	return s.repository.FindSubLayananByID(id)
}

func (s *subLayananService) GetSubLayananByLayananID(layananDesaID uint) ([]domain.SubLayananDesa, error) {
	return s.repository.FindSubLayananByLayananID(layananDesaID)
}

func (s *subLayananService) UpdateSubLayanan(id uint, req *types.SubLayananRequest) (*domain.SubLayananDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	subLayanan, err := s.repository.FindSubLayananByID(id)
	if err != nil {
		return nil, err
	}

	// Validasi layanan desa exists jika diubah
	if req.LayananDesaID != subLayanan.LayananDesaID {
		_, err := s.layananRepository.FindLayananByID(req.LayananDesaID)
		if err != nil {
			return nil, errors.New("layanan desa tidak ditemukan")
		}
	}

	subLayanan.LayananDesaID = req.LayananDesaID
	subLayanan.Nama = req.Nama
	subLayanan.Persyaratan = req.Persyaratan
	subLayanan.Template = req.Template
	subLayanan.Aktif = req.Aktif

	err = s.repository.UpdateSubLayanan(subLayanan)
	return subLayanan, err
}

func (s *subLayananService) DeleteSubLayanan(id uint) error {
	subLayanan, err := s.repository.FindSubLayananByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteSubLayanan(subLayanan)
}

// ==================== PENGAJUAN LAYANAN SERVICE ====================

type PengajuanLayananService interface {
	CreatePengajuan(req *types.PengajuanRequest) (*domain.PengajuanLayanan, error)
	GetAllPengajuan(
		search string,
		status string,
		nik string,
		layananDesaID uint,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.PengajuanLayanan, error)
	GetPengajuanByID(id uint) (*domain.PengajuanLayanan, error)
	GetPengajuanByNIK(nik string) ([]domain.PengajuanLayanan, error)
	GetPengajuanByNomorSurat(nomorSurat string) (*domain.PengajuanLayanan, error)
	UpdatePengajuan(id uint, req *types.PengajuanUpdateRequest) (*domain.PengajuanLayanan, error)
	UpdateStatusPengajuan(id uint, req *types.UpdateStatusRequest, updatedBy uint) (*domain.PengajuanLayanan, error)
	ApprovePengajuan(id uint, req *types.ApproveRequest, approvedBy uint) (*domain.PengajuanLayanan, error)
	RejectPengajuan(id uint, req *types.RejectRequest, rejectedBy uint) (*domain.PengajuanLayanan, error)
	DeletePengajuan(id uint) error
	GetStatisticsByStatus() (map[string]int64, error)
}

type pengajuanLayananService struct {
	repository         PengajuanLayananRepository
	layananRepository  LayananRepository
	subLayananRepo     SubLayananRepository
	riwayatRepository  RiwayatPengajuanRepository
	validate           *validator.Validate
}

func NewPengajuanLayananService(
	repository PengajuanLayananRepository,
	layananRepository LayananRepository,
	subLayananRepo SubLayananRepository,
	riwayatRepository RiwayatPengajuanRepository,
) PengajuanLayananService {
	return &pengajuanLayananService{
		repository:        repository,
		layananRepository: layananRepository,
		subLayananRepo:    subLayananRepo,
		riwayatRepository: riwayatRepository,
		validate:          validator.New(),
	}
}

func (s *pengajuanLayananService) CreatePengajuan(req *types.PengajuanRequest) (*domain.PengajuanLayanan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Validasi layanan desa exists
	_, err := s.layananRepository.FindLayananByID(req.LayananDesaID)
	if err != nil {
		return nil, errors.New("layanan desa tidak ditemukan")
	}

	// Validasi sub layanan exists (jika ada)
	if req.SubLayananID != nil && *req.SubLayananID > 0 {
		_, err := s.subLayananRepo.FindSubLayananByID(*req.SubLayananID)
		if err != nil {
			return nil, errors.New("sub layanan tidak ditemukan")
		}
	}

	pengajuan := domain.PengajuanLayanan{
		LayananDesaID: req.LayananDesaID,
		SubLayananID:  req.SubLayananID,
		NamaLengkap:   req.NamaLengkap,
		NIK:           req.NIK,
		TempatLahir:   req.TempatLahir,
		TanggalLahir:  req.TanggalLahir,
		JenisKelamin:  req.JenisKelamin,
		Alamat:        req.Alamat,
		RT:            req.RT,
		RW:            req.RW,
		NoTelp:        req.NoTelp,
		Email:         req.Email,
		Keperluan:     req.Keperluan,
		Catatan:       req.Catatan,
		Dokumen:       req.Dokumen,
		Status:        "Pending",
	}

	err = s.repository.CreatePengajuan(&pengajuan)
	if err != nil {
		return nil, err
	}

	// Buat riwayat pengajuan
	riwayat := domain.RiwayatPengajuan{
		PengajuanID: pengajuan.ID,
		StatusLama:  "",
		StatusBaru:  "Pending",
		Keterangan:  "Pengajuan baru dibuat",
	}
	s.riwayatRepository.CreateRiwayat(&riwayat)

	return &pengajuan, nil
}

func (s *pengajuanLayananService) GetAllPengajuan(
	search string,
	status string,
	nik string,
	layananDesaID uint,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.PengajuanLayanan, error) {
	return s.repository.FindPengajuan(search, status, nik, layananDesaID, page, limit, sortBy, sortOrder)
}

func (s *pengajuanLayananService) GetPengajuanByID(id uint) (*domain.PengajuanLayanan, error) {
	return s.repository.FindPengajuanByID(id)
}

func (s *pengajuanLayananService) GetPengajuanByNIK(nik string) ([]domain.PengajuanLayanan, error) {
	return s.repository.FindPengajuanByNIK(nik)
}

func (s *pengajuanLayananService) GetPengajuanByNomorSurat(nomorSurat string) (*domain.PengajuanLayanan, error) {
	return s.repository.FindPengajuanByNomorSurat(nomorSurat)
}

func (s *pengajuanLayananService) UpdatePengajuan(id uint, req *types.PengajuanUpdateRequest) (*domain.PengajuanLayanan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	pengajuan, err := s.repository.FindPengajuanByID(id)
	if err != nil {
		return nil, err
	}

	// Hanya bisa update jika status masih Pending
	if pengajuan.Status != "Pending" {
		return nil, errors.New("pengajuan tidak dapat diubah karena sudah diproses")
	}

	pengajuan.NamaLengkap = req.NamaLengkap
	pengajuan.NIK = req.NIK
	pengajuan.TempatLahir = req.TempatLahir
	pengajuan.TanggalLahir = req.TanggalLahir
	pengajuan.JenisKelamin = req.JenisKelamin
	pengajuan.Alamat = req.Alamat
	pengajuan.RT = req.RT
	pengajuan.RW = req.RW
	pengajuan.NoTelp = req.NoTelp
	pengajuan.Email = req.Email
	pengajuan.Keperluan = req.Keperluan
	pengajuan.Catatan = req.Catatan
	pengajuan.Dokumen = req.Dokumen

	err = s.repository.UpdatePengajuan(pengajuan)
	return pengajuan, err
}

func (s *pengajuanLayananService) UpdateStatusPengajuan(id uint, req *types.UpdateStatusRequest, updatedBy uint) (*domain.PengajuanLayanan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	pengajuan, err := s.repository.FindPengajuanByID(id)
	if err != nil {
		return nil, err
	}

	statusLama := pengajuan.Status
	pengajuan.Status = req.Status
	pengajuan.Keterangan = req.Keterangan

	now := time.Now()
	pengajuan.ProcessedBy = &updatedBy
	pengajuan.ProcessedAt = &now

	err = s.repository.UpdatePengajuan(pengajuan)
	if err != nil {
		return nil, err
	}

	// Buat riwayat perubahan status
	riwayat := domain.RiwayatPengajuan{
		PengajuanID: pengajuan.ID,
		StatusLama:  statusLama,
		StatusBaru:  req.Status,
		Keterangan:  req.Keterangan,
		UpdatedBy:   &updatedBy,
	}
	s.riwayatRepository.CreateRiwayat(&riwayat)

	return pengajuan, nil
}

func (s *pengajuanLayananService) ApprovePengajuan(id uint, req *types.ApproveRequest, approvedBy uint) (*domain.PengajuanLayanan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	pengajuan, err := s.repository.FindPengajuanByID(id)
	if err != nil {
		return nil, err
	}

	statusLama := pengajuan.Status
	pengajuan.Status = "Approved"
	pengajuan.NomorSurat = req.NomorSurat
	pengajuan.TanggalSurat = &req.TanggalSurat
	pengajuan.Keterangan = req.Keterangan

	now := time.Now()
	pengajuan.ProcessedBy = &approvedBy
	pengajuan.ProcessedAt = &now

	err = s.repository.UpdatePengajuan(pengajuan)
	if err != nil {
		return nil, err
	}

	// Buat riwayat persetujuan
	riwayat := domain.RiwayatPengajuan{
		PengajuanID: pengajuan.ID,
		StatusLama:  statusLama,
		StatusBaru:  "Approved",
		Keterangan:  "Pengajuan disetujui. " + req.Keterangan,
		UpdatedBy:   &approvedBy,
	}
	s.riwayatRepository.CreateRiwayat(&riwayat)

	return pengajuan, nil
}

func (s *pengajuanLayananService) RejectPengajuan(id uint, req *types.RejectRequest, rejectedBy uint) (*domain.PengajuanLayanan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	pengajuan, err := s.repository.FindPengajuanByID(id)
	if err != nil {
		return nil, err
	}

	statusLama := pengajuan.Status
	pengajuan.Status = "Rejected"
	pengajuan.Keterangan = req.Keterangan

	now := time.Now()
	pengajuan.ProcessedBy = &rejectedBy
	pengajuan.ProcessedAt = &now

	err = s.repository.UpdatePengajuan(pengajuan)
	if err != nil {
		return nil, err
	}

	// Buat riwayat penolakan
	riwayat := domain.RiwayatPengajuan{
		PengajuanID: pengajuan.ID,
		StatusLama:  statusLama,
		StatusBaru:  "Rejected",
		Keterangan:  "Pengajuan ditolak. " + req.Keterangan,
		UpdatedBy:   &rejectedBy,
	}
	s.riwayatRepository.CreateRiwayat(&riwayat)

	return pengajuan, nil
}

func (s *pengajuanLayananService) DeletePengajuan(id uint) error {
	pengajuan, err := s.repository.FindPengajuanByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeletePengajuan(pengajuan)
}

func (s *pengajuanLayananService) GetStatisticsByStatus() (map[string]int64, error) {
	statuses := []string{"Pending", "Processing", "Approved", "Rejected", "Completed"}
	statistics := make(map[string]int64)

	for _, status := range statuses {
		count, err := s.repository.CountByStatus(status)
		if err != nil {
			return nil, err
		}
		statistics[status] = count
	}

	return statistics, nil
}

// ==================== RIWAYAT PENGAJUAN SERVICE ====================

type RiwayatPengajuanService interface {
	GetRiwayatByPengajuanID(pengajuanID uint) ([]domain.RiwayatPengajuan, error)
	GetAllRiwayat(
		pengajuanID uint,
		page int,
		limit int,
	) ([]domain.RiwayatPengajuan, error)
}

type riwayatPengajuanService struct {
	repository RiwayatPengajuanRepository
}

func NewRiwayatPengajuanService(repository RiwayatPengajuanRepository) RiwayatPengajuanService {
	return &riwayatPengajuanService{
		repository: repository,
	}
}

func (s *riwayatPengajuanService) GetRiwayatByPengajuanID(pengajuanID uint) ([]domain.RiwayatPengajuan, error) {
	return s.repository.FindRiwayatByPengajuanID(pengajuanID)
}

func (s *riwayatPengajuanService) GetAllRiwayat(
	pengajuanID uint,
	page int,
	limit int,
) ([]domain.RiwayatPengajuan, error) {
	return s.repository.FindRiwayat(pengajuanID, page, limit)
}
