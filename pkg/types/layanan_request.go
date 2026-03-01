package types

import "time"

type LayananRequest struct {
	ServiceName   string `json:"service_name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Category      string `json:"category" binding:"required"`
	Image         string `json:"image" binding:"required"`
	EstimatedTime string `json:"estimated_time" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Slug          string `json:"slug" binding:"required"`
	Cost          string `json:"cost" binding:"required"`
}

type PengajuanLayananResponse struct {
	ID           uint       `json:"id"`
	NamaLengkap  string     `json:"nama_lengkap"`
	NIK          string     `json:"nik"`
	LayananDesa  string     `json:"layanan_desa"`
	SubLayanan   string     `json:"sub_layanan,omitempty"`
	Status       string     `json:"status"`
	NomorSurat   string     `json:"nomor_surat,omitempty"`
	TanggalSurat *time.Time `json:"tanggal_surat,omitempty"`
	Keperluan    string     `json:"keperluan"`
	CreatedAt    time.Time  `json:"created_at"`
}

type PengajuanLayananRequest struct {
	JenisLayanan  string   `json:"jenis_layanan" binding:"required"` // ID SubLayanan
	NamaLengkap   string   `json:"nama_lengkap" binding:"required"`
	NIK           string   `json:"nik" binding:"required,len=16"`
	TempatLahir   string   `json:"tempat_lahir" binding:"required"`
	TanggalLahir  string   `json:"tanggal_lahir" binding:"required"`
	JenisKelamin  string   `json:"jenis_kelamin" binding:"required,oneof=L P"`
	Alamat        string   `json:"alamat" binding:"required"`
	RT            string   `json:"rt" binding:"required"`
	RW            string   `json:"rw" binding:"required"`
	NoTelp        string   `json:"no_telp" binding:"required"`
	Email         string   `json:"email" binding:"omitempty,email"`
	Keperluan     string   `json:"keperluan" binding:"required"`
	Catatan       string   `json:"catatan"`
	Files         []string `json:"files"` // File paths setelah upload
}

