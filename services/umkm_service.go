package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type UmkmService interface {
	CreateUmkm(req *requests.UmkmRequest) (*models.Umkm, error)
	GetAllUmkm() ([]models.Umkm, error)
	GetUmkmByID(id uint) (*models.Umkm, error)
	GetUmkmBySlug(slug string) (*models.Umkm, error)
	UpdateUmkm(id uint, req *requests.UmkmRequest) (*models.Umkm, error)
	DeleteUmkm(id uint) error
}

type umkmService struct {
	repository repositories.UmkmRepository
	validate   *validator.Validate
}

func NewUmkmService(repository repositories.UmkmRepository) UmkmService {
	return &umkmService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *umkmService) CreateUmkm(req *requests.UmkmRequest) (*models.Umkm, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	umkm := models.Umkm{
		NamaUsaha: req.NamaUsaha,
		Kategori:  req.Kategori,
		Deskripsi: req.Deskripsi,
		Gambar:    req.Gambar,
		Pemilik:   req.Pemilik,
		Status:    req.Status,
		Slug:      req.Slug,
	}

	err := s.repository.CreateUmkm(&umkm)
	return &umkm, err
}

func (s *umkmService) GetAllUmkm() ([]models.Umkm, error) {
	return s.repository.FindUmkm()
}

func (s *umkmService) GetUmkmByID(id uint) (*models.Umkm, error) {
	return s.repository.FindUmkmByID(id)
}

func (s *umkmService) GetUmkmBySlug(slug string) (*models.Umkm, error) {
	return s.repository.FindUmkmBySlug(slug)
}

func (s *umkmService) UpdateUmkm(id uint, req *requests.UmkmRequest) (*models.Umkm, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	umkm, err := s.repository.FindUmkmByID(id)
	if err != nil {
		return nil, err
	}

	umkm.NamaUsaha = req.NamaUsaha
	umkm.Kategori = req.Kategori
	umkm.Deskripsi = req.Deskripsi
	umkm.Gambar = req.Gambar
	umkm.Pemilik = req.Pemilik
	umkm.Status = req.Status
	umkm.Slug = req.Slug

	err = s.repository.UpdateUmkm(umkm)
	return umkm, err
}

func (s *umkmService) DeleteUmkm(id uint) error {
	umkm, err := s.repository.FindUmkmByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteUmkm(umkm)
}
