package requests

type LayananKesehatanRequest struct {
	NamaProgram  string `json:"nama_program" validate:"required"`
	Deskripsi    string `json:"deskripsi"`
	JenisProgram string `json:"jenis_program" validate:"required"`
	FasilitasID  string  `json:"fasilitas_id"`
	Jadwal       string `json:"jadwal"`
}

type FasilitasKesehatanRequest struct {
	NamaFasilitas   string `json:"nama_fasilitas" validate:"required"`
	Alamat          string `json:"alamat"`
	PenanggungJawab string `json:"penanggung_jawab"`
	NoTelepon       string `json:"no_telepon"`
	JamOperasional  string `json:"jam_operasional"`
	Jenis           string `json:"jenis" validate:"required"`
}
