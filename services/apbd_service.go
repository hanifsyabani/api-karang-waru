package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type ApbdService interface {
	CreateApbd(req *requests.ApbdRequest) (*models.APBDDesa, error)
	GetAllApbd() ([]models.APBDDesa, error)
	GetApbdByID(id uint) (*models.APBDDesa, error)
	UpdateApbd(id uint, req *requests.ApbdRequest) (*models.APBDDesa, error)
	DeleteApbd(id uint) error
}

type apbdService struct {
	repository repositories.ApbdDesaRepository
	validate   *validator.Validate
}

func NewApbdService(repository repositories.ApbdDesaRepository) ApbdService {
	return &apbdService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *apbdService) CreateApbd(req *requests.ApbdRequest) (*models.APBDDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	totalBelanja := req.BelanjaPenyelenggaraanPemerintahan + req.BelanjaPembangunan + req.BelanjaPembinaanKemasyarakatan + req.BelanjaPemberdayaanMasyarakat + req.BelanjaTakTerduga
	totalPendapatan := req.PendapatanAsliDesa + req.Transfer + req.PendapatanLain
	surplus := totalPendapatan - totalBelanja


	apbd := models.APBDDesa{
		Tahun:                              req.Tahun,
		PendapatanAsliDesa:                 req.PendapatanAsliDesa,
		Transfer:                           req.Transfer,
		PendapatanLain:                     req.PendapatanLain,
		BelanjaPenyelenggaraanPemerintahan: req.BelanjaPenyelenggaraanPemerintahan,
		BelanjaPembangunan:                 req.BelanjaPembangunan,
		BelanjaPembinaanKemasyarakatan:     req.BelanjaPembinaanKemasyarakatan,
		BelanjaPemberdayaanMasyarakat:      req.BelanjaPemberdayaanMasyarakat,
		BelanjaTakTerduga:                  req.BelanjaTakTerduga,
		PenerimaanPembiayaan:               req.PenerimaanPembiayaan,
		PengeluaranPembiayaan:              req.PengeluaranPembiayaan,

		
		TotalPendapatan:                    totalPendapatan,
		TotalBelanja:                       totalBelanja,
		SurplusDefisit:                     surplus,
		Status:                             req.Status,
		Keterangan:                         req.Keterangan,
		FileLampiran:                       req.FileLampiran,
	}

	err := s.repository.CreateApbd(&apbd)
	return &apbd, err
}

func (s *apbdService) GetAllApbd() ([]models.APBDDesa, error) {
	return s.repository.FindApbd()
}

func (s *apbdService) GetApbdByID(id uint) (*models.APBDDesa, error) {
	return s.repository.FindApbdByID(id)
}

func (s *apbdService) UpdateApbd(id uint, req *requests.ApbdRequest) (*models.APBDDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	apbd, err := s.repository.FindApbdByID(id)
	if err != nil {
		return nil, err
	}

	apbd.Tahun = req.Tahun
	apbd.PendapatanAsliDesa = req.PendapatanAsliDesa
	apbd.Transfer = req.Transfer
	apbd.PendapatanLain = req.PendapatanLain
	apbd.BelanjaPenyelenggaraanPemerintahan = req.BelanjaPenyelenggaraanPemerintahan
	apbd.BelanjaPembangunan = req.BelanjaPembangunan
	apbd.BelanjaPembinaanKemasyarakatan = req.BelanjaPembinaanKemasyarakatan
	apbd.BelanjaPemberdayaanMasyarakat = req.BelanjaPemberdayaanMasyarakat
	apbd.BelanjaTakTerduga = req.BelanjaTakTerduga
	apbd.PenerimaanPembiayaan = req.PenerimaanPembiayaan
	apbd.PengeluaranPembiayaan = req.PengeluaranPembiayaan


	// calc totalBelanja
	totalBelanja := apbd.BelanjaPenyelenggaraanPemerintahan + apbd.BelanjaPembangunan + apbd.BelanjaPembinaanKemasyarakatan + apbd.BelanjaPemberdayaanMasyarakat + apbd.BelanjaTakTerduga
	totalPendapatan := apbd.PendapatanAsliDesa + apbd.Transfer + apbd.PendapatanLain
	surplus := totalPendapatan - totalBelanja

	apbd.TotalPendapatan = totalPendapatan
	apbd.TotalBelanja = totalBelanja
	apbd.SurplusDefisit = surplus
	apbd.Status = req.Status
	apbd.Keterangan = req.Keterangan
	apbd.FileLampiran = req.FileLampiran

	err = s.repository.UpdateApbd(apbd)
	return apbd, err
}

func (s *apbdService) DeleteApbd(id uint) error {
	apbd, err := s.repository.FindApbdByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteApbd(apbd)
}
