package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type SejarahService interface {
	CreateSejarah(req *requests.SejarahRequest) (*models.Sejarah, error)
	GetSejarah() (models.Sejarah, error)
	UpdateSejarah(req *requests.SejarahRequest) (*models.Sejarah, error)
	DeleteSejarah() error
}

type sejarahService struct {
	repository repositories.SejarahRepository
	validate   *validator.Validate
}

func NewSejarahService(repository repositories.SejarahRepository) SejarahService {
	return &sejarahService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *sejarahService) CreateSejarah(req *requests.SejarahRequest) (*models.Sejarah, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	sejarah := models.Sejarah{
		ProfilDesaID: req.ProfilDesaID,
		Body:         req.Body,
	}

	err := s.repository.CreateSejarah(&sejarah)
	return &sejarah, err
}

func (s *sejarahService) GetSejarah() (models.Sejarah, error) {
	return s.repository.FindSejarah()
}

func (s *sejarahService) UpdateSejarah(req *requests.SejarahRequest) (*models.Sejarah, error) {
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
