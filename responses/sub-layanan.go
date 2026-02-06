package responses

import (
	"api-karang-waru/models"
	"time"
)

// SubLayananResponse
type SubLayananResponse struct {
	ID            uint      `json:"id"`
	LayananDesaID uint      `json:"layanan_desa_id"`
	Nama          string    `json:"nama"`
	Persyaratan   string    `json:"persyaratan"`
	Template      string    `json:"template"`
	Aktif         bool      `json:"aktif"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func SubLayananResponseFromModel(model *models.SubLayananDesa) SubLayananResponse {
	return SubLayananResponse{
		ID:            model.ID,
		LayananDesaID: model.LayananDesaID,
		Nama:          model.Nama,
		Persyaratan:   model.Persyaratan,
		Template:      model.Template,
		Aktif:         model.Aktif,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// PengajuanResponse
type PengajuanResponse struct {
	ID            uint               `json:"id"`
	LayananDesa   *LayananResponse   `json:"layanan_desa,omitempty"`
	SubLayanan    *SubLayananResponse `json:"sub_layanan,omitempty"`
	NamaLengkap   string             `json:"nama_lengkap"`
	NIK           string             `json:"nik"`
	TempatLahir   string             `json:"tempat_lahir"`
	TanggalLahir  string             `json:"tanggal_lahir"`
	JenisKelamin  string             `json:"jenis_kelamin"`
	Alamat        string             `json:"alamat"`
	RT            string             `json:"rt"`
	RW            string             `json:"rw"`
	NoTelp        string             `json:"no_telp"`
	Email         string             `json:"email"`
	Keperluan     string             `json:"keperluan"`
	Catatan       string             `json:"catatan"`
	Dokumen       string             `json:"dokumen"`
	Status        string             `json:"status"`
	Keterangan    string             `json:"keterangan"`
	NomorSurat    string             `json:"nomor_surat"`
	TanggalSurat  *time.Time         `json:"tanggal_surat"`
	ProcessedBy   *uint              `json:"processed_by"`
	ProcessedAt   *time.Time         `json:"processed_at"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

func PengajuanResponseFromModel(model *models.PengajuanLayanan) PengajuanResponse {
	resp := PengajuanResponse{
		ID:           model.ID,
		NamaLengkap:  model.NamaLengkap,
		NIK:          model.NIK,
		TempatLahir:  model.TempatLahir,
		TanggalLahir: model.TanggalLahir,
		JenisKelamin: model.JenisKelamin,
		Alamat:       model.Alamat,
		RT:           model.RT,
		RW:           model.RW,
		NoTelp:       model.NoTelp,
		Email:        model.Email,
		Keperluan:    model.Keperluan,
		Catatan:      model.Catatan,
		Dokumen:      model.Dokumen,
		Status:       model.Status,
		Keterangan:   model.Keterangan,
		NomorSurat:   model.NomorSurat,
		TanggalSurat: model.TanggalSurat,
		ProcessedBy:  model.ProcessedBy,
		ProcessedAt:  model.ProcessedAt,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}

	// Include relations if loaded
	if model.LayananDesa.ID != 0 {
		layananResp := LayananResponseFromModel(&model.LayananDesa)
		resp.LayananDesa = &layananResp
	}

	if model.SubLayanan != nil && model.SubLayanan.ID != 0 {
		subLayananResp := SubLayananResponseFromModel(model.SubLayanan)
		resp.SubLayanan = &subLayananResp
	}

	return resp
}

// RiwayatPengajuanResponse
type RiwayatPengajuanResponse struct {
	ID          uint      `json:"id"`
	PengajuanID uint      `json:"pengajuan_id"`
	StatusLama  string    `json:"status_lama"`
	StatusBaru  string    `json:"status_baru"`
	Keterangan  string    `json:"keterangan"`
	UpdatedBy   *uint     `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func RiwayatPengajuanResponseFromModel(model *models.RiwayatPengajuan) RiwayatPengajuanResponse {
	return RiwayatPengajuanResponse{
		ID:          model.ID,
		PengajuanID: model.PengajuanID,
		StatusLama:  model.StatusLama,
		StatusBaru:  model.StatusBaru,
		Keterangan:  model.Keterangan,
		UpdatedBy:   model.UpdatedBy,
		CreatedAt:   model.CreatedAt,
	}
}