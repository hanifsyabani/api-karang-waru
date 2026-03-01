package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type VisiMisiService interface {
	CreateVisiMisi(req *types.VisiMisiRequest) (*domain.VisiMisi, error)
	GetVisiMisi() (domain.VisiMisi, error)
	UpdateVisiMisi(req *types.VisiMisiRequest) (*domain.VisiMisi, error)
	DeleteVisiMisi() error
}

type visiMisiService struct {
	repository VisiMisiRepository
	validate   *validator.Validate
}

func NewVisiMisiService(repository VisiMisiRepository) VisiMisiService {
	return &visiMisiService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *visiMisiService) CreateVisiMisi(req *types.VisiMisiRequest) (*domain.VisiMisi, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	visiMisi := domain.VisiMisi{
		ProfilDesaID: req.ProfilDesaID,
		Visi:         req.Visi,
		Misi:         req.Misi,
	}

	err := s.repository.CreateVisiMisi(&visiMisi)
	return &visiMisi, err
}

func (s *visiMisiService) GetVisiMisi() (domain.VisiMisi, error) {
	return s.repository.FindVisiMisi()
}

func (s *visiMisiService) UpdateVisiMisi(req *types.VisiMisiRequest) (*domain.VisiMisi, error) {
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

