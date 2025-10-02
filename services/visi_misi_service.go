package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type VisiMisiService interface {
	CreateVisiMisi(req *requests.VisiMisiRequest) (*models.VisiMisi, error)
	GetVisiMisi() (models.VisiMisi, error)
	UpdateVisiMisi(req *requests.VisiMisiRequest) (*models.VisiMisi, error)
	DeleteVisiMisi() error
}

type visiMisiService struct {
	repository repositories.VisiMisiRepository
	validate   *validator.Validate
}

func NewVisiMisiService(repository repositories.VisiMisiRepository) VisiMisiService {
	return &visiMisiService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *visiMisiService) CreateVisiMisi(req *requests.VisiMisiRequest) (*models.VisiMisi, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	visiMisi := models.VisiMisi{
		ProfilDesaID: req.ProfilDesaID,
		Visi:         req.Visi,
		Misi:         req.Misi,
	}

	err := s.repository.CreateVisiMisi(&visiMisi)
	return &visiMisi, err
}

func (s *visiMisiService) GetVisiMisi() (models.VisiMisi, error) {
	return s.repository.FindVisiMisi()
}

func (s *visiMisiService) UpdateVisiMisi(req *requests.VisiMisiRequest) (*models.VisiMisi, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	visiMisi, err := s.repository.FindVisiMisi()
	if err != nil {
		return nil, err
	}

	visiMisi.ProfilDesaID = req.ProfilDesaID
	visiMisi.Visi = req.Visi
	visiMisi.Misi = req.Misi

	err = s.repository.UpdateVisiMisi(&visiMisi)
	return &visiMisi, err
}

func (s *visiMisiService) DeleteVisiMisi() error {
	return s.repository.DeleteVisiMisi()
}
