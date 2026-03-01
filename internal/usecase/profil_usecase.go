package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"

	"github.com/go-playground/validator/v10"
)

type ProfilService interface {
	CreateProfil(req *types.ProfilDesaRequest) (*domain.ProfilDesa, error)

	GetProfil() (domain.ProfilDesa, error)

	UpdateProfil(req *types.ProfilDesaRequest) (*domain.ProfilDesa, error)

	DeleteProfil() error
}

type profilService struct {
	repository ProfilRepository
	validate   *validator.Validate
}

func NewProfilService(repository ProfilRepository) ProfilService {
	return &profilService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *profilService) CreateProfil(req *types.ProfilDesaRequest) (*domain.ProfilDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	profil := domain.ProfilDesa{
		Alamat:          req.Alamat,
		Kecamatan:       req.Kecamatan,
		Kabupaten:       req.Kabupaten,
		Provinsi:        req.Provinsi,
		KodePos:         req.KodePos,
		JumlahPenduduk:  req.JumlahPenduduk,
		JumlahLaki:      req.JumlahLaki,
		JumlahPerempuan: req.JumlahPerempuan,
		JumlahKK:        req.JumlahKK,

		TahunPembentukan: req.TahunPembentukan,
		Telepon:          req.Telepon,
		Email:            req.Email,
	}

	err := s.repository.CreateProfil(&profil)
	return &profil, err
}

func (s *profilService) GetProfil() (domain.ProfilDesa, error) {
	return s.repository.FindProfil()
}

func (s *profilService) UpdateProfil(req *types.ProfilDesaRequest) (*domain.ProfilDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	profil, err := s.repository.FindProfil()
	if err != nil {
		return nil, err
	}

	profil.Alamat = req.Alamat
	profil.Kecamatan = req.Kecamatan
	profil.Kabupaten = req.Kabupaten
	profil.Provinsi = req.Provinsi
	profil.KodePos = req.KodePos
	profil.JumlahPenduduk = req.JumlahPenduduk
	profil.JumlahLaki = req.JumlahLaki
	profil.JumlahPerempuan = req.JumlahPerempuan
	profil.JumlahKK = req.JumlahKK

	profil.TahunPembentukan = req.TahunPembentukan
	profil.Telepon = req.Telepon
	profil.Email = req.Email

	err = s.repository.UpdateProfil(&profil)
	return &profil, err
}

func (s *profilService) DeleteProfil() error {
	return s.repository.DeleteProfil()
}



