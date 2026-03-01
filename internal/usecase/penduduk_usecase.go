package usecase

import (
	customErr "api-karang-waru/pkg/errors"
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"
	"github.com/go-playground/validator/v10"
	"time"
)

type PendudukService interface {
	CreatePenduduk(req *types.PendudukRequest) (*domain.Penduduk, error)
	GetAllPenduduk(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.Penduduk, error)
	CountPenduduk() (
		total int64,
		lakiLaki int64,
		perempuan int64,
		kartuKeluarga int64,
		err error,
	)
	GetPendudukByID(id uint) (*domain.Penduduk, error)
	UpdatePenduduk(id uint, req *types.PendudukRequest) (*domain.Penduduk, error)
	DeletePenduduk(id uint) error
}

type pendudukService struct {
	repository PendudukRepository
	validate   *validator.Validate
}

func NewPendudukService(repository PendudukRepository) PendudukService {
	return &pendudukService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *pendudukService) CreatePenduduk(req *types.PendudukRequest) (*domain.Penduduk, error) {
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

	penduduk := domain.Penduduk{
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
) ([]domain.Penduduk, error) {
	return s.repository.FindPenduduk(search, page, limit, sortBy, sortOrder)
}

func (s *pendudukService) GetPendudukByID(id uint) (*domain.Penduduk, error) {
	return s.repository.FindPendudukByID(id)
}

func (s *pendudukService) CountPenduduk() (
	total int64,
	lakiLaki int64,
	perempuan int64,
	kartuKeluarga int64,
	err error,
) {
	return s.repository.CountPenduduk()
}

func (s *pendudukService) UpdatePenduduk(id uint, req *types.PendudukRequest) (*domain.Penduduk, error) {
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

