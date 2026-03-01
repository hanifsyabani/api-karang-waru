package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type DemografisService interface {
	CreateDemografis(req *types.DemografisRequest) (*domain.Demografis, error)
	GetDemografis() (domain.Demografis, error)
	UpdateDemografis(req *types.DemografisRequest) (*domain.Demografis, error)
	DeleteDemografis() error
}

type demografisService struct {
	repository DemografisRepository
	validate   *validator.Validate
}

func NewDemografisService(repository DemografisRepository) DemografisService {
	return &demografisService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *demografisService) CreateDemografis(req *types.DemografisRequest) (*domain.Demografis, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	demografis := domain.Demografis{
		ProfilDesaID:     req.ProfilDesaID,
		Balita:           req.Balita,
		Anak:             req.Anak,
		Dewasa:           req.Dewasa,
		Lansia:           req.Lansia,
		Pertanian:        req.Pertanian,
		Perdagangan:      req.Perdagangan,
		Jasa:             req.Jasa,
		Industri:         req.Industri,
		Sekolah:          req.Sekolah,
		Puskesmas:        req.Puskesmas,
		Masjid:           req.Masjid,
		PasarTradisional: req.PasarTradisional,
		PosKeamanan:      req.PosKeamanan,
		BalaiDesa:        req.BalaiDesa,
	}

	err := s.repository.CreateDemografis(&demografis)
	return &demografis, err
}

func (s *demografisService) GetDemografis() (domain.Demografis, error) {
	return s.repository.FindDemografis()
}

func (s *demografisService) UpdateDemografis(req *types.DemografisRequest) (*domain.Demografis, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	demografis, err := s.repository.FindDemografis()
	if err != nil {
		return nil, err
	}

	demografis.ProfilDesaID = req.ProfilDesaID
	demografis.Balita = req.Balita
	demografis.Anak = req.Anak
	demografis.Dewasa = req.Dewasa
	demografis.Lansia = req.Lansia
	demografis.Pertanian = req.Pertanian
	demografis.Perdagangan = req.Perdagangan
	demografis.Jasa = req.Jasa
	demografis.Industri = req.Industri
	demografis.Sekolah = req.Sekolah
	demografis.Puskesmas = req.Puskesmas
	demografis.Masjid = req.Masjid
	demografis.PasarTradisional = req.PasarTradisional
	demografis.PosKeamanan = req.PosKeamanan
	demografis.BalaiDesa = req.BalaiDesa

	err = s.repository.UpdateDemografis(&demografis)
	return &demografis, err
}

func (s *demografisService) DeleteDemografis() error {
	return s.repository.DeleteDemografis()
}

