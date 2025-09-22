package responses

import "api-karang-waru/models"

type ProfilDesaResponse struct {
	Alamat           string  `json:"alamat"`
	Kecamatan        string  `json:"kecamatan"`
	Kabupaten        string  `json:"kabupaten"`
	Provinsi         string  `json:"provinsi"`
	KodePos          string  `json:"kode_pos"`
	JumlahPenduduk   string     `json:"jumlah_penduduk"`
	JumlahLaki       string     `json:"jumlah_laki"`
	JumlahPerempuan  string     `json:"jumlah_perempuan"`
	JumlahKK         string     `json:"jumlah_kk"`

	TahunPembentukan string     `json:"tahun_pembentukan"`
	Telepon          string  `json:"telepon"`
	Email            string  `json:"email"`
}

func ProfilResponseFromModel(profil *models.ProfilDesa) ProfilDesaResponse {
	return ProfilDesaResponse{
		Alamat:           profil.Alamat,
		Kecamatan:        profil.Kecamatan,
		Kabupaten:        profil.Kabupaten,
		Provinsi:         profil.Provinsi,
		KodePos:          profil.KodePos,
		JumlahPenduduk:   profil.JumlahPenduduk,
		JumlahLaki:       profil.JumlahLaki,
		JumlahPerempuan:  profil.JumlahPerempuan,
		JumlahKK:         profil.JumlahKK,

		TahunPembentukan: profil.TahunPembentukan,
		Telepon:          profil.Telepon,
		Email:            profil.Email,
	}
}
