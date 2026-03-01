package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type UmkmService interface {
	CreateUmkm(req *types.UmkmRequest) (*domain.Umkm, error)
	GetAllUmkm(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
		status string,
	) ([]domain.Umkm, error)
	GetUmkmByID(id uint) (*domain.Umkm, error)
	GetCountStatus() (verified int64, unverified int64, err error)
	GetUmkmBySlug(slug string) (*domain.Umkm, error)
	UpdateUmkm(id uint, req *types.UmkmRequest) (*domain.Umkm, error)
	DeleteUmkm(id uint) error
}

type umkmService struct {
	repository UmkmRepository
	validate   *validator.Validate
}

func NewUmkmService(repository UmkmRepository) UmkmService {
	return &umkmService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *umkmService) CreateUmkm(req *types.UmkmRequest) (*domain.Umkm, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	umkm := domain.Umkm{
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

func (s *umkmService) GetAllUmkm(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
	status string,
) ([]domain.Umkm, error) {
	return s.repository.FindUmkm(
		search,
		page,
		limit,
		sortBy,
		sortOrder,
		status,
	)
}

func (s *umkmService) GetUmkmByID(id uint) (*domain.Umkm, error) {
	return s.repository.FindUmkmByID(id)
}

func (s *umkmService) GetUmkmBySlug(slug string) (*domain.Umkm, error) {
	return s.repository.FindUmkmBySlug(slug)
}
func (s *umkmService) GetCountStatus() (verified int64, unverified int64, err error) {
	return s.repository.CountStatus()
}


func (s *umkmService) UpdateUmkm(id uint, req *types.UmkmRequest) (*domain.Umkm, error) {
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

