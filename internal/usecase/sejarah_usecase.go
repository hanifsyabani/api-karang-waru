package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type SejarahService interface {
	CreateSejarah(req *types.SejarahRequest) (*domain.Sejarah, error)
	GetSejarah() (domain.Sejarah, error)
	UpdateSejarah(req *types.SejarahRequest) (*domain.Sejarah, error)
	DeleteSejarah() error
}

type sejarahService struct {
	repository SejarahRepository
	validate   *validator.Validate
}

func NewSejarahService(repository SejarahRepository) SejarahService {
	return &sejarahService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *sejarahService) CreateSejarah(req *types.SejarahRequest) (*domain.Sejarah, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	sejarah := domain.Sejarah{
		ProfilDesaID: req.ProfilDesaID,
		Body:         req.Body,
	}

	err := s.repository.CreateSejarah(&sejarah)
	return &sejarah, err
}

func (s *sejarahService) GetSejarah() (domain.Sejarah, error) {
	return s.repository.FindSejarah()
}

func (s *sejarahService) UpdateSejarah(req *types.SejarahRequest) (*domain.Sejarah, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	sejarah, err := s.repository.FindSejarah()
	if err != nil {
		return nil, err
	}

	sejarah.ProfilDesaID = req.ProfilDesaID
	sejarah.Body = req.Body

	err = s.repository.UpdateSejarah(&sejarah)
	return &sejarah, err
}

func (s *sejarahService) DeleteSejarah() error {
	return s.repository.DeleteSejarah()
}

