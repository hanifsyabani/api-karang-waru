package responses

import "api-karang-waru/models"

type ProfilDesaResponse struct {
	Alamat           string  `json:"alamat"`
	Kecamatan        string  `json:"kecamatan"`
	Kabupaten        string  `json:"kabupaten"`
	Provinsi         string  `json:"provinsi"`
	KodePos          string  `json:"kode_pos"`
	JumlahPenduduk   int     `json:"jumlah_penduduk"`
	JumlahLaki       int     `json:"jumlah_laki"`
	JumlahPerempuan  int     `json:"jumlah_perempuan"`
	JumlahKK         int     `json:"jumlah_kk"`
	LuasWilayahKm2   float64 `json:"luas_wilayah_km2"`
	LuasWilayahHa    int     `json:"luas_wilayah_ha"`
	TahunPembentukan int     `json:"tahun_pembentukan"`
	Telepon          string  `json:"telepon"`
	Email            string  `json:"email"`
	JamPelayanan     string  `json:"jam_pelayanan"`
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
		LuasWilayahKm2:   profil.LuasWilayahKm2,
		LuasWilayahHa:    profil.LuasWilayahHa,
		TahunPembentukan: profil.TahunPembentukan,
		Telepon:          profil.Telepon,
		Email:            profil.Email,
		JamPelayanan:     profil.JamPelayanan,
	}
}
