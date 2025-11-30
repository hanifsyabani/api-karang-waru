package responses

import (
	"api-karang-waru/models"
)

type LayananKesehatanResponse struct {
	ID           uint   `json:"id"`
	NamaProgram  string `json:"nama_program"`
	Deskripsi    string `json:"deskripsi"`
	JenisProgram string `json:"jenis_program"`

	FasilitasID *uint  `json:"fasilitas_id"`
	Fasilitas   string `json:"fasilitas,omitempty"`

	Jadwal string `json:"jadwal"`
	Status string `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type FasilitasKesehatanResponse struct {
	ID              uint   `json:"id"`
	NamaFasilitas   string `json:"nama_fasilitas"`
	Alamat          string `json:"alamat"`
	NoTelepon       string `json:"no_telepon"`
	PenanggungJawab string `json:"penanggung_jawab"`
	JamOperasional  string `json:"jam_operasional"`
	Jenis           string `json:"jenis"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func LayananKesehatanResponseFromModel(l *models.LayananKesehatan) LayananKesehatanResponse {
	resp := LayananKesehatanResponse{
		ID:           l.ID,
		NamaProgram:  l.NamaProgram,
		Deskripsi:    l.Deskripsi,
		JenisProgram: l.JenisProgram,
		FasilitasID:  l.FasilitasID,
		Jadwal:       l.Jadwal,
		CreatedAt:    l.CreatedAt.Format("2006-01-02"),
		UpdatedAt:    l.UpdatedAt.Format("2006-01-02"),
	}

	// Jika relasi fasilitas dimuat, masukkan nama fasilitas
	if l.Fasilitas.NamaFasilitas != "" {
		resp.Fasilitas = l.Fasilitas.NamaFasilitas
	}

	return resp
}
func FasilitasKesehatanResponseFromModel(l *models.FasilitasKesehatan) FasilitasKesehatanResponse {
	resp := FasilitasKesehatanResponse{
		ID:              l.ID,
		NamaFasilitas:   l.NamaFasilitas,
		Alamat:          l.Alamat,
		NoTelepon:       l.NoTelepon,
		PenanggungJawab: l.PenanggungJawab,
		JamOperasional:  l.JamOperasional,
		Jenis:           l.Jenis,
		CreatedAt:       l.CreatedAt.Format("2006-01-02"),
		UpdatedAt:       l.UpdatedAt.Format("2006-01-02"),
	}

	return resp
}
