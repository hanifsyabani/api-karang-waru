package types

import "time"

// SubLayananRequest - Request untuk create/update sub layanan
type SubLayananRequest struct {
	LayananDesaID uint   `json:"layanan_desa_id" validate:"required"`
	Nama          string `json:"nama" validate:"required"`
	Persyaratan   string `json:"persyaratan"`
	Template      string `json:"template"`
	Aktif         bool   `json:"aktif"`
}

// PengajuanRequest - Request untuk create pengajuan
type PengajuanRequest struct {
	LayananDesaID uint   `json:"layanan_desa_id" validate:"required"`
	SubLayananID  *uint  `json:"sub_layanan_id"`
	NamaLengkap   string `json:"nama_lengkap" validate:"required"`
	NIK           string `json:"nik" validate:"required,len=16"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin" validate:"oneof=L P"`
	Alamat        string `json:"alamat"`
	RT            string `json:"rt"`
	RW            string `json:"rw"`
	NoTelp        string `json:"no_telp" validate:"required"`
	Email         string `json:"email" validate:"omitempty,email"`
	Keperluan     string `json:"keperluan" validate:"required"`
	Catatan       string `json:"catatan"`
	Dokumen       string `json:"dokumen"`
}

// PengajuanUpdateRequest - Request untuk update pengajuan
type PengajuanUpdateRequest struct {
	NamaLengkap  string `json:"nama_lengkap" validate:"required"`
	NIK          string `json:"nik" validate:"required,len=16"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin" validate:"oneof=L P"`
	Alamat       string `json:"alamat"`
	RT           string `json:"rt"`
	RW           string `json:"rw"`
	NoTelp       string `json:"no_telp" validate:"required"`
	Email        string `json:"email" validate:"omitempty,email"`
	Keperluan    string `json:"keperluan" validate:"required"`
	Catatan      string `json:"catatan"`
	Dokumen      string `json:"dokumen"`
}

// UpdateStatusRequest - Request untuk update status pengajuan
type UpdateStatusRequest struct {
	Status     string `json:"status" validate:"required,oneof=Pending Processing Approved Rejected Completed"`
	Keterangan string `json:"keterangan"`
}

// ApproveRequest - Request untuk approve pengajuan
type ApproveRequest struct {
	NomorSurat   string    `json:"nomor_surat" validate:"required"`
	TanggalSurat time.Time `json:"tanggal_surat" validate:"required"`
	Keterangan   string    `json:"keterangan"`
}

// RejectRequest - Request untuk reject pengajuan
type RejectRequest struct {
	Keterangan string `json:"keterangan" validate:"required"`
}
