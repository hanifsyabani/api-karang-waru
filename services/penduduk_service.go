package services

import (
	customErr "api-karang-waru/errors"
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"
	"github.com/go-playground/validator/v10"
	"time"
)

type PendudukService interface {
	CreatePenduduk(req *requests.PendudukRequest) (*models.Penduduk, error)
	GetAllPenduduk(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.Penduduk, error)
	CountPenduduk() (int64, error)
	GetPendudukByID(id uint) (*models.Penduduk, error)
	UpdatePenduduk(id uint, req *requests.PendudukRequest) (*models.Penduduk, error)
	DeletePenduduk(id uint) error
}

type pendudukService struct {
	repository repositories.PendudukRepository
	validate   *validator.Validate
}

func NewPendudukService(repository repositories.PendudukRepository) PendudukService {
	return &pendudukService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *pendudukService) CreatePenduduk(req *requests.PendudukRequest) (*models.Penduduk, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}
	parsedDate, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return nil, err
	}

	exists, err := s.repository.IsNIKExists(req.NIK)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, customErr.NewNIKAlreadyExistsError(req.NIK)
	}

	penduduk := models.Penduduk{
		NIK:                req.NIK,
		NoKK:               req.NoKK,
		NamaLengkap:        req.NamaLengkap,
		JenisKelamin:       req.JenisKelamin,
		TempatLahir:        req.TempatLahir,
		TanggalLahir:       parsedDate,
		Dusun:              req.Dusun,
		Desa:               req.Desa,
		Kecamatan:          req.Kecamatan,
		Kabupaten:          req.Kabupaten,
		Provinsi:           req.Provinsi,
		Alamat:             req.Alamat,
		RT:                 req.RT,
		RW:                 req.RW,
		Agama:              req.Agama,
		StatusPerkawinan:   req.StatusPerkawinan,
		Pekerjaan:          req.Pekerjaan,
		PendidikanTerakhir: req.PendidikanTerakhir,
		Kewarganegaraan:    req.Kewarganegaraan,
		StatusKependudukan: req.StatusKependudukan,
		Keterangan:         req.Keterangan,
	}

	if err := s.repository.CreatePenduduk(&penduduk); err != nil {
		return nil, err
	}
	return &penduduk, err
}

func (s *pendudukService) GetAllPenduduk(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.Penduduk, error) {
	return s.repository.FindPenduduk(search, page, limit, sortBy, sortOrder)
}

func (s *pendudukService) GetPendudukByID(id uint) (*models.Penduduk, error) {
	return s.repository.FindPendudukByID(id)
}

func (s *pendudukService) CountPenduduk() (int64, error) {
	return s.repository.CountPenduduk()
}

func (s *pendudukService) UpdatePenduduk(id uint, req *requests.PendudukRequest) (*models.Penduduk, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	penduduk, err := s.repository.FindPendudukByID(id)
	if err != nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return nil, err
	}

	penduduk.NIK = req.NIK
	penduduk.NoKK = req.NoKK
	penduduk.NamaLengkap = req.NamaLengkap
	penduduk.JenisKelamin = req.JenisKelamin
	penduduk.TempatLahir = req.TempatLahir
	penduduk.TanggalLahir = parsedDate
	penduduk.Dusun = req.Dusun
	penduduk.Desa = req.Desa
	penduduk.Kecamatan = req.Kecamatan
	penduduk.Kabupaten = req.Kabupaten
	penduduk.Provinsi = req.Provinsi
	penduduk.Alamat = req.Alamat
	penduduk.RT = req.RT
	penduduk.RW = req.RW
	penduduk.Agama = req.Agama
	penduduk.StatusPerkawinan = req.StatusPerkawinan
	penduduk.Pekerjaan = req.Pekerjaan
	penduduk.PendidikanTerakhir = req.PendidikanTerakhir
	penduduk.Kewarganegaraan = req.Kewarganegaraan
	penduduk.StatusKependudukan = req.StatusKependudukan
	penduduk.Keterangan = req.Keterangan

	err = s.repository.UpdatePenduduk(penduduk)
	return penduduk, err
}

func (s *pendudukService) DeletePenduduk(id uint) error {
	penduduk, err := s.repository.FindPendudukByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeletePenduduk(penduduk)
}
