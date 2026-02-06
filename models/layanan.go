package models

import (
	"time"

	"gorm.io/gorm"
)

// LayananDesa - Model untuk layanan utama desa
type LayananDesa struct {
	ID            uint   `gorm:"primaryKey;column:id"`
	ServiceName   string `gorm:"column:service_name;type:VARCHAR(255);not null"` // Service title, e.g., "Layanan Surat Menyurat"
	Description   string `gorm:"column:description;type:TEXT"`                   // Detailed explanation, requirements, and workflow
	Category      string `gorm:"column:category;type:VARCHAR(100)"`              // e.g., "Administrasi", "Kependudukan", "Sosial"
	Image         string `gorm:"column:image;type:VARCHAR(255)"`                 // Path to an icon or image for the service
	Slug          string `gorm:"column:slug;type:VARCHAR(255);uniqueIndex"`      // For clean URLs (e.g., /services/surat-menyurat)
	Status        string `gorm:"column:status;type:VARCHAR(50);default:'Draft'"` // "Published", "Draft"
	
	EstimatedTime string `gorm:"column:estimated_time;type:VARCHAR(100)"`        // e.g., "1-3 Hari Kerja"
	Cost          string `gorm:"column:cost;type:VARCHAR(100)"`                  // e.g., "Gratis"

	// Relations
	SubLayanan []SubLayananDesa `gorm:"foreignKey:LayananDesaID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (LayananDesa) TableName() string {
	return "village_services"
}

// SubLayananDesa - Model untuk sub-layanan (jenis-jenis surat)
type SubLayananDesa struct {
	ID             uint   `gorm:"primaryKey;column:id"`
	LayananDesaID  uint   `gorm:"column:layanan_desa_id;not null;index"` // Foreign key ke LayananDesa
	Nama           string `gorm:"column:nama;type:VARCHAR(255);not null"` // e.g., "Surat Keterangan Domisili"
	Persyaratan    string `gorm:"column:persyaratan;type:TEXT"`           // JSON array: ["KTP Asli", "KK Asli"]
	Template       string `gorm:"column:template;type:VARCHAR(255)"`      // Path to document template
	Aktif          bool   `gorm:"column:aktif;type:BOOLEAN;default:true"` // Apakah sub-layanan aktif

	// Relations
	Pengajuan []PengajuanLayanan `gorm:"foreignKey:SubLayananID;constraint:OnDelete:SET NULL"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (SubLayananDesa) TableName() string {
	return "village_sub_services"
}

// PengajuanLayanan - Model untuk pengajuan layanan dari warga
type PengajuanLayanan struct {
	ID            uint   `gorm:"primaryKey;column:id"`
	LayananDesaID uint   `gorm:"column:layanan_desa_id;not null;index"` // Foreign key ke LayananDesa
	SubLayananID  *uint  `gorm:"column:sub_layanan_id;index"`           // Foreign key ke SubLayananDesa (nullable)
	
	// Data Pemohon
	NamaLengkap   string `gorm:"column:nama_lengkap;type:VARCHAR(255);not null"`
	NIK           string `gorm:"column:nik;type:VARCHAR(16);not null;index"`
	TempatLahir   string `gorm:"column:tempat_lahir;type:VARCHAR(100)"`
	TanggalLahir  string `gorm:"column:tanggal_lahir;type:DATE"`
	JenisKelamin  string `gorm:"column:jenis_kelamin;type:VARCHAR(1)"` // "L" atau "P"
	
	// Alamat
	Alamat        string `gorm:"column:alamat;type:TEXT"`
	RT            string `gorm:"column:rt;type:VARCHAR(3)"`
	RW            string `gorm:"column:rw;type:VARCHAR(3)"`
	
	// Kontak
	NoTelp        string `gorm:"column:no_telp;type:VARCHAR(20);not null"`
	Email         string `gorm:"column:email;type:VARCHAR(100)"`
	
	// Keperluan
	Keperluan     string `gorm:"column:keperluan;type:TEXT;not null"`
	Catatan       string `gorm:"column:catatan;type:TEXT"`
	
	// Dokumen Upload (JSON array of file paths)
	Dokumen       string `gorm:"column:dokumen;type:TEXT"` // JSON: ["path/to/file1.pdf", "path/to/file2.jpg"]
	
	// Status Pengajuan
	Status        string `gorm:"column:status;type:VARCHAR(50);default:'Pending'"` // "Pending", "Processing", "Approved", "Rejected", "Completed"
	Keterangan    string `gorm:"column:keterangan;type:TEXT"`                      // Keterangan dari admin
	
	// Nomor Surat (jika sudah disetujui)
	NomorSurat    string `gorm:"column:nomor_surat;type:VARCHAR(100);uniqueIndex"`
	TanggalSurat  *time.Time `gorm:"column:tanggal_surat;type:DATE"`
	
	// Admin yang memproses
	ProcessedBy   *uint      `gorm:"column:processed_by"` // Foreign key ke User/Admin
	ProcessedAt   *time.Time `gorm:"column:processed_at;type:TIMESTAMP"`

	// Relations
	LayananDesa   LayananDesa    `gorm:"foreignKey:LayananDesaID"`
	SubLayanan    *SubLayananDesa `gorm:"foreignKey:SubLayananID"`
	Riwayat       []RiwayatPengajuan `gorm:"foreignKey:PengajuanID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (PengajuanLayanan) TableName() string {
	return "service_submissions"
}

// RiwayatPengajuan - Model untuk tracking riwayat status pengajuan
type RiwayatPengajuan struct {
	ID           uint   `gorm:"primaryKey;column:id"`
	PengajuanID  uint   `gorm:"column:pengajuan_id;not null;index"` // Foreign key ke PengajuanLayanan
	StatusLama   string `gorm:"column:status_lama;type:VARCHAR(50)"`
	StatusBaru   string `gorm:"column:status_baru;type:VARCHAR(50);not null"`
	Keterangan   string `gorm:"column:keterangan;type:TEXT"`
	UpdatedBy    *uint  `gorm:"column:updated_by"` // Foreign key ke User/Admin
	
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
}

func (RiwayatPengajuan) TableName() string {
	return "submission_history"
}

// Struct untuk response API (DTO)


