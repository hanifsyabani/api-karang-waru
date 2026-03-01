package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"
	"time"

	"github.com/go-playground/validator/v10"
)

type BeritaService interface {
	CreateBerita(req *types.BeritaRequest)(*domain.Berita, error)
	GetAllBerita(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	)([]domain.Berita, error)
	GetBeritaByID(id uint)(*domain.Berita, error)
	GetCountBeritaByCategory() (map[string]int64, error)
	GetBeritaBySlug(slug string)(*domain.Berita, error)
	UpdateBerita(id uint, req *types.BeritaRequest)(*domain.Berita, error)
	DeleteBerita(id uint) error
}

type beritaService struct {
	repository BeritaRepository
	validate *validator.Validate
}

func NewBeritaService(repository BeritaRepository) BeritaService{
	return &beritaService{
		repository: repository,
		validate: validator.New(),
	}
}

func (s *beritaService) CreateBerita(req *types.BeritaRequest) (*domain.Berita,error) {
	if err := s.validate.Struct(req);err !=nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}

	berita := domain.Berita{
		Title: req.Title,
		Slug: req.Slug,
		Category: req.Category,
		Content: req.Content,
		Image: req.Image,
		Writer: req.Writer,
		Date: parsedDate,
		Status: req.Status,
	}

	err = s.repository.CreateBerita(&berita)
	return &berita, err
}

func (s *beritaService) GetAllBerita(search string , page int, limit int, sortBy string, sortOrder string) ([]domain.Berita, error) {
	return s.repository.FindBerita(
		search,
		page,
		limit,
		sortBy,
		sortOrder,
	)
}

func (s *beritaService) GetCountBeritaByCategory() (map[string]int64, error) {
	return s.repository.CountBeritaByCategory()
}

func (s *beritaService) GetBeritaByID(id uint) (*domain.Berita, error) {
	return s.repository.FindBeritaByID(id)
}

func (s *beritaService) GetBeritaBySlug(slug string) (*domain.Berita, error) {
	return  s.repository.FindBeritaBySlug(slug)
}

func (s *beritaService) UpdateBerita(id uint, req *types.BeritaRequest)(*domain.Berita,error){
	if err := s.validate.Struct(req);err !=nil{
		return nil, err
	}

	berita,err := s.repository.FindBeritaByID(id)
	if err != nil{
		return nil, err
	}


	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}

	berita.Title = req.Title
	berita.Slug = req.Slug
	berita.Category = req.Category
	berita.Content = req.Content
	berita.Image = req.Image
	berita.Writer = req.Writer
	berita.Date = parsedDate
	berita.Status = req.Status

	err = s.repository.UpdateBerita(berita)
	return berita, err
}

func (s *beritaService) DeleteBerita(id uint) error {
	berita, err := s.repository.FindBeritaByID(id)
	if err != nil{
		return err
	}

	return s.repository.DeleteBerita(berita)
}
