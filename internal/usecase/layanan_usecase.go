package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type LayananService interface {
	CreateLayanan(req *types.LayananRequest) (*domain.LayananDesa, error)
	GetAllLayanan(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.LayananDesa, error)
	GetLayananByID(id uint) (*domain.LayananDesa, error)
	GetLayananBySlug(slug string) (*domain.LayananDesa, error)
	UpdateLayanan(id uint, req *types.LayananRequest) (*domain.LayananDesa, error)
	DeleteLayanan(id uint) error
}

type layananService struct {
	repository LayananRepository
	validate   *validator.Validate
}

func NewLayananService(repository LayananRepository) LayananService {
	return &layananService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *layananService) CreateLayanan(req *types.LayananRequest) (*domain.LayananDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	layanan := domain.LayananDesa{
		ServiceName:   req.ServiceName,
		Description:   req.Description,
		Category:      req.Category,
		Image:         req.Image,
		EstimatedTime: req.EstimatedTime,
		Cost:          req.Cost,
		Status:        req.Status,
		Slug:          req.Slug,
	}

	err := s.repository.CreateLayanan(&layanan)
	return &layanan, err
}

func (s *layananService) GetAllLayanan(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.LayananDesa, error) {
	return s.repository.FindLayanan(search, page, limit, sortBy, sortOrder)
}

func (s *layananService) GetLayananByID(id uint) (*domain.LayananDesa, error) {
	return s.repository.FindLayananByID(id)
}

func (s *layananService) GetLayananBySlug(slug string) (*domain.LayananDesa, error) {
	return s.repository.FindLayananBySlug(slug)
}

func (s *layananService) UpdateLayanan(id uint, req *types.LayananRequest) (*domain.LayananDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	layanan, err := s.repository.FindLayananByID(id)
	if err != nil {
		return nil, err
	}

	layanan.ServiceName = req.ServiceName
	layanan.Description = req.Description
	layanan.Category = req.Category
	layanan.Image = req.Image
	layanan.EstimatedTime = req.EstimatedTime
	layanan.Cost = req.Cost
	layanan.Status = req.Status
	layanan.Slug = req.Slug

	err = s.repository.UpdateLayanan(layanan)
	return layanan, err
}

func (s *layananService) DeleteLayanan(id uint) error {
	layanan, err := s.repository.FindLayananByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteLayanan(layanan)
}

