package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type DemografisService interface {
	CreateDemografis(req *requests.DemografisRequest) (*models.Demografis, error)
	GetDemografis() (models.Demografis, error)
	UpdateDemografis(req *requests.DemografisRequest) (*models.Demografis, error)
	DeleteDemografis() error
}

type demografisService struct {
	repository repositories.DemografisRepository
	validate   *validator.Validate
}

func NewDemografisService(repository repositories.DemografisRepository) DemografisService {
	return &demografisService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *demografisService) CreateDemografis(req *requests.DemografisRequest) (*models.Demografis, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	demografis := models.Demografis{
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

func (s *demografisService) GetDemografis() (models.Demografis, error) {
	return s.repository.FindDemografis()
}

func (s *demografisService) UpdateDemografis(req *requests.DemografisRequest) (*models.Demografis, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	demografis, err := s.repository.FindDemografis()
	if err != nil {
		return nil, err
	}

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