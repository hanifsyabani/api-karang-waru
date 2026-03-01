package types

import (
	"api-karang-waru/internal/domain"
)

// PendudukResponse merepresentasikan data yang dikirim ke client
type PendudukResponse struct {
	ID                 uint   `json:"id"`
	NIK                string `json:"nik"`
	NoKK               string `json:"no_kk"`
	NamaLengkap        string `json:"nama_lengkap"`
	JenisKelamin       string `json:"jenis_kelamin"`
	TempatLahir        string `json:"tempat_lahir"`
	TanggalLahir       string `json:"tanggal_lahir"`
	Alamat             string `json:"alamat"`
	RT                 string `json:"rt"`
	RW                 string `json:"rw"`
	Dusun              string `json:"dusun"`
	Desa               string `json:"desa"`
	Kecamatan          string `json:"kecamatan"`
	Kabupaten          string `json:"kabupaten"`
	Provinsi           string `json:"provinsi"`
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	Pekerjaan          string `json:"pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	Kewarganegaraan    string `json:"kewarganegaraan"`
	StatusKependudukan string `json:"status_kependudukan"`
	Keterangan          string `json:"keterangan"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// PendudukResponseFromModel mengubah model Penduduk menjadi response JSON
func PendudukResponseFromModel(penduduk *domain.Penduduk) PendudukResponse {
	return PendudukResponse{
		ID:                 penduduk.ID,
		NIK:                penduduk.NIK,
		NoKK:               penduduk.NoKK,
		NamaLengkap:        penduduk.NamaLengkap,
		JenisKelamin:       penduduk.JenisKelamin,
		TempatLahir:        penduduk.TempatLahir,
		TanggalLahir:       penduduk.TanggalLahir.Format("2006-01-02"),
		Alamat:             penduduk.Alamat,
		RT:                 penduduk.RT,
		RW:                 penduduk.RW,
		Dusun:              penduduk.Dusun,
		Desa:               penduduk.Desa,
		Kecamatan:          penduduk.Kecamatan,
		Kabupaten:          penduduk.Kabupaten,
		Provinsi:           penduduk.Provinsi,
		Agama:              penduduk.Agama,
		StatusPerkawinan:   penduduk.StatusPerkawinan,
		Pekerjaan:          penduduk.Pekerjaan,
		PendidikanTerakhir: penduduk.PendidikanTerakhir,
		Kewarganegaraan:    penduduk.Kewarganegaraan,
		StatusKependudukan: penduduk.StatusKependudukan,
		Keterangan:          penduduk.Keterangan,
		CreatedAt:          penduduk.CreatedAt.Format("2006-01-02"),
		UpdatedAt:          penduduk.UpdatedAt.Format("2006-01-02"),
	}
}

