package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"

	"github.com/go-playground/validator/v10"
)

type ProfilService interface {
	CreateProfil(req *requests.ProfilDesaRequest) (*models.ProfilDesa, error)
	GetProfil() (models.ProfilDesa, error)
	UpdateProfil(req *requests.ProfilDesaRequest) (*models.ProfilDesa, error)
	DeleteProfil() error
}

type profilService struct {
	repository repositories.ProfilRepository
	validate   *validator.Validate
}

func NewProfilService(repository repositories.ProfilRepository) ProfilService {
	return &profilService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *profilService) CreateProfil(req *requests.ProfilDesaRequest) (*models.ProfilDesa, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	profil := models.ProfilDesa{
		Alamat:           req.Alamat,
		Kecamatan:        req.Kecamatan,
		Kabupaten:        req.Kabupaten,
		Provinsi:         req.Provinsi,
		KodePos:          req.KodePos,
		JumlahPenduduk:   req.JumlahPenduduk,
		JumlahLaki:       req.JumlahLaki,
		JumlahPerempuan:  req.JumlahPerempuan,
		JumlahKK:         req.JumlahKK,
		LuasWilayahKm2:   req.LuasWilayahKm2,
		LuasWilayahHa:    req.LuasWilayahHa,
		TahunPembentukan: req.TahunPembentukan,
		Telepon:          req.Telepon,
		Email:            req.Email,
		JamPelayanan:     req.JamPelayanan,
	}

	err := s.repository.CreateProfil(&profil)
	return &profil, err
}

func (s *profilService) GetProfil() (models.ProfilDesa, error) {
	return s.repository.FindProfil()
}

func (s *profilService) UpdateProfil(req *requests.ProfilDesaRequest) (*models.ProfilDesa, error) {
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
	profil.LuasWilayahKm2 = req.LuasWilayahKm2
	profil.LuasWilayahHa = req.LuasWilayahHa
	profil.TahunPembentukan = req.TahunPembentukan
	profil.Telepon = req.Telepon
	profil.Email = req.Email
	profil.JamPelayanan = req.JamPelayanan

	err = s.repository.UpdateProfil(&profil)
	return &profil, err
}

func (s *profilService) DeleteProfil() error {
	return s.repository.DeleteProfil()
}
