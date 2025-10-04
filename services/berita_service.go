package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"
	"time"

	"github.com/go-playground/validator/v10"
)

type BeritaService interface {
	CreateBerita(req *requests.BeritaRequest)(*models.Berita, error)
	GetAllBerita()([]models.Berita, error)
	GetBeritaByID(id uint)(*models.Berita, error)
	GetBeritaBySlug(slug string)(*models.Berita, error)
	UpdateBerita(id uint, req *requests.BeritaRequest)(*models.Berita, error)
	DeleteBerita(id uint) error
}

type beritaService struct {
	repository repositories.BeritaRepository
	validate *validator.Validate
}

func NewBeritaService(repository repositories.BeritaRepository) BeritaService{
	return &beritaService{
		repository: repository,
		validate: validator.New(),
	}
}

func (s *beritaService) CreateBerita(req *requests.BeritaRequest) (*models.Berita,error) {
	if err := s.validate.Struct(req);err !=nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}

	berita := models.Berita{
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

func (s *beritaService) GetAllBerita() ([]models.Berita, error) {
	return s.repository.FindBerita()
}

func (s *beritaService) GetBeritaByID(id uint) (*models.Berita, error) {
	return s.repository.FindBeritaByID(id)
}

func (s *beritaService) GetBeritaBySlug(slug string) (*models.Berita, error) {
	return  s.repository.FindBeritaBySlug(slug)
}

func (s *beritaService) UpdateBerita(id uint, req *requests.BeritaRequest)(*models.Berita,error){
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