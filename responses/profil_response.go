package responses

import "api-karang-waru/models"

type ProfilDesaResponse struct {
	ID              uint   `json:"id"`
	Alamat          string `json:"alamat"`
	Kecamatan       string `json:"kecamatan"`
	Kabupaten       string `json:"kabupaten"`
	Provinsi        string `json:"provinsi"`
	KodePos         string `json:"kode_pos"`
	JumlahPenduduk  string `json:"jumlah_penduduk"`
	JumlahLaki      string `json:"jumlah_laki"`
	JumlahPerempuan string `json:"jumlah_perempuan"`
	JumlahKK        string `json:"jumlah_kk"`

	TahunPembentukan string `json:"tahun_pembentukan"`
	Telepon          string `json:"telepon"`
	Email            string `json:"email"`
}

func ProfilResponseFromModel(profil *models.ProfilDesa) ProfilDesaResponse {
	return ProfilDesaResponse{
		ID:              profil.ID,
		Alamat:          profil.Alamat,
		Kecamatan:       profil.Kecamatan,
		Kabupaten:       profil.Kabupaten,
		Provinsi:        profil.Provinsi,
		KodePos:         profil.KodePos,
		JumlahPenduduk:  profil.JumlahPenduduk,
		JumlahLaki:      profil.JumlahLaki,
		JumlahPerempuan: profil.JumlahPerempuan,
		JumlahKK:        profil.JumlahKK,

		TahunPembentukan: profil.TahunPembentukan,
		Telepon:          profil.Telepon,
		Email:            profil.Email,
	}
}

type DemografisResponse struct {
	ProfilDesaID     uint   `json:"profil_desa_id"`
	Balita           string `json:"balita"`
	Anak             string `json:"anak"`
	Dewasa           string `json:"dewasa"`
	Lansia           string `json:"lansia"`
	Pertanian        string `json:"pertanian"`
	Perdagangan      string `json:"perdagangan"`
	Jasa             string `json:"jasa"`
	Industri         string `json:"industri"`
	Sekolah          string `json:"sekolah"`
	Puskesmas        string `json:"puskesmas"`
	Masjid           string `json:"masjid"`
	PasarTradisional string `json:"pasar_tradisional"`
	PosKeamanan      string `json:"pos_keamanan"`
	BalaiDesa        string `json:"balai_desa"`
}

func DemografisResponseFromModel(profil *models.Demografis) DemografisResponse {
	return DemografisResponse{
		ProfilDesaID:     profil.ProfilDesaID,
		Balita:           profil.Balita,
		Anak:             profil.Anak,
		Dewasa:           profil.Dewasa,
		Lansia:           profil.Lansia,
		Pertanian:        profil.Pertanian,
		Perdagangan:      profil.Perdagangan,
		Jasa:             profil.Jasa,
		Industri:         profil.Industri,
		Sekolah:          profil.Sekolah,
		Puskesmas:        profil.Puskesmas,
		Masjid:           profil.Masjid,
		PasarTradisional: profil.PasarTradisional,
		PosKeamanan:      profil.PosKeamanan,
		BalaiDesa:        profil.BalaiDesa,
	}
}
