package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type KesehatanService interface {
	CreateLayanan(req *types.LayananKesehatanRequest) (*domain.LayananKesehatan, error)
	GetAllLayanan() ([]domain.LayananKesehatan, error)
	GetLayananByID(id uint) (*domain.LayananKesehatan, error)
	UpdateLayanan(id uint, req *types.LayananKesehatanRequest) (*domain.LayananKesehatan, error)
	DeleteLayanan(id uint) error

	CreateFasilitas(req *types.FasilitasKesehatanRequest) (*domain.FasilitasKesehatan, error)
	GetAllFasilitas() ([]domain.FasilitasKesehatan, error)
	GetFasilitasByID(id uint) (*domain.FasilitasKesehatan, error)
	UpdateFasilitas(id uint, req *types.FasilitasKesehatanRequest) (*domain.FasilitasKesehatan, error)
	DeleteFasilitas(id uint) error
}

type kesehatanService struct {
	repository KesehatanRepository
	validate   *validator.Validate
}

func NewKesehatanService(repository KesehatanRepository) KesehatanService {
	return &kesehatanService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *kesehatanService) CreateLayanan(req *types.LayananKesehatanRequest) (*domain.LayananKesehatan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}



	fid, err := strconv.Atoi(req.FasilitasID)
	if err != nil {
		return nil, fmt.Errorf("fasilitas_id harus berupa angka")
	}

    fidUint := uint(fid)

	layanan := domain.LayananKesehatan{
		NamaProgram:  req.NamaProgram,
		Deskripsi:    req.Deskripsi,
		JenisProgram: req.JenisProgram,
		FasilitasID:  &fidUint,
		Jadwal:       req.Jadwal,
	}

	err = s.repository.CreateLayanan(&layanan)
	return &layanan, err
}

func (s *kesehatanService) GetAllLayanan() ([]domain.LayananKesehatan, error) {
	return s.repository.FindAllLayanan()
}

func (s *kesehatanService) GetLayananByID(id uint) (*domain.LayananKesehatan, error) {
	return s.repository.FindLayananByID(id)
}

func (s *kesehatanService) UpdateLayanan(id uint, req *types.LayananKesehatanRequest) (*domain.LayananKesehatan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	layanan, err := s.repository.FindLayananByID(id)
	if err != nil {
		return nil, err
	}

    fid, err := strconv.Atoi(req.FasilitasID)
    if err != nil {
        return nil, fmt.Errorf("fasilitas_id harus berupa angka")
    }

    fidUint := uint(fid)

	layanan.NamaProgram = req.NamaProgram
	layanan.Deskripsi = req.Deskripsi
	layanan.JenisProgram = req.JenisProgram
	layanan.FasilitasID = &fidUint 
	layanan.Jadwal = req.Jadwal

	err = s.repository.UpdateLayanan(layanan)
	return layanan, err
}

func (s *kesehatanService) DeleteLayanan(id uint) error {
	layanan, err := s.repository.FindLayananByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteLayanan(layanan)
}

func (s *kesehatanService) CreateFasilitas(req *types.FasilitasKesehatanRequest) (*domain.FasilitasKesehatan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	fasilitas := domain.FasilitasKesehatan{
		NamaFasilitas:   req.NamaFasilitas,
		Alamat:          req.Alamat,
		PenanggungJawab: req.PenanggungJawab,
		NoTelepon:       req.NoTelepon,
		JamOperasional:  req.JamOperasional,
		Jenis:           req.Jenis,
	}

	err := s.repository.CreateFasilitasKesehatan(&fasilitas)
	return &fasilitas, err
}

func (s *kesehatanService) GetAllFasilitas() ([]domain.FasilitasKesehatan, error) {
	return s.repository.FindAllFasilitasKesehatan()
}

func (s *kesehatanService) GetFasilitasByID(id uint) (*domain.FasilitasKesehatan, error) {
	return s.repository.FindFasilitasKesehatanByID(id)
}

func (s *kesehatanService) UpdateFasilitas(id uint, req *types.FasilitasKesehatanRequest) (*domain.FasilitasKesehatan, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	fasilitas, err := s.repository.FindFasilitasKesehatanByID(id)
	if err != nil {
		return nil, err
	}

	fasilitas.NamaFasilitas = req.NamaFasilitas
	fasilitas.Alamat = req.Alamat
	fasilitas.PenanggungJawab = req.PenanggungJawab
	fasilitas.NoTelepon = req.NoTelepon
	fasilitas.JamOperasional = req.JamOperasional
	fasilitas.Jenis = req.Jenis

	err = s.repository.UpdateFasilitasKesehatan(fasilitas)
	return fasilitas, err
}

func (s *kesehatanService) DeleteFasilitas(id uint) error {
	fasilitas, err := s.repository.FindFasilitasKesehatanByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteFasilitasKesehatan(fasilitas)
}

