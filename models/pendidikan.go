package models

import (
	"time"

	"gorm.io/gorm"
)

// ==============================
// 1. Data Lembaga Pendidikan
// ==============================
type LembagaPendidikan struct {
	ID uint `gorm:"primaryKey;column:id"`

	NamaSekolah       string  `gorm:"column:nama_sekolah;type:VARCHAR(150);not null"`
	JenjangPendidikan string  `gorm:"column:jenjang_pendidikan;type:VARCHAR(50);not null"` // PAUD, TK, SD, SMP, SMA, Perguruan Tinggi, dll
	Alamat            string  `gorm:"column:alamat;type:TEXT;not null"`
	Latitude          float64 `gorm:"column:latitude;type:DECIMAL(10,8)"`
	Longitude         float64 `gorm:"column:longitude;type:DECIMAL(11,8)"`

	JumlahSiswa int    `gorm:"column:jumlah_siswa;default:0"`
	JumlahGuru  int    `gorm:"column:jumlah_guru;default:0"`
	JumlahStaf  int    `gorm:"column:jumlah_staf;default:0"`
	Kontak      string `gorm:"column:kontak;type:VARCHAR(100)"` // Telepon atau email

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (LembagaPendidikan) TableName() string {
	return "lembaga_pendidikan"
}

// ==============================
// 2. Statistik Pendidikan Penduduk
// ==============================
type StatistikPendidikan struct {
	ID uint `gorm:"primaryKey;column:id"`

	Tahun           string `gorm:"column:tahun;type:VARCHAR(4);not null;index"`
	TidakSekolah    int    `gorm:"column:tidak_sekolah;default:0"`
	SD              int    `gorm:"column:sd;default:0"`
	SMP             int    `gorm:"column:smp;default:0"`
	SMA             int    `gorm:"column:sma;default:0"`
	PerguruanTinggi int    `gorm:"column:perguruan_tinggi;default:0"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (StatistikPendidikan) TableName() string {
	return "statistik_pendidikan"
}

// ==============================
// 3. Program Pendidikan Desa
// ==============================
type ProgramPendidikan struct {
	ID uint `gorm:"primaryKey;column:id"`

	NamaProgram    string    `gorm:"column:nama_program;type:VARCHAR(150);not null"`
	Deskripsi      string    `gorm:"column:deskripsi;type:TEXT"`
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai"`
	Status         string    `gorm:"column:status;type:VARCHAR(50);default:'Aktif'"` // Aktif / Selesai

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (ProgramPendidikan) TableName() string {
	return "program_pendidikan"
}

// ==============================
// 4. Capaian Pendidikan
// ==============================
type CapaianPendidikan struct {
	ID uint `gorm:"primaryKey;column:id"`

	Tahun              string  `gorm:"column:tahun;type:VARCHAR(4);not null;index"`
	JumlahLulusan      int     `gorm:"column:jumlah_lulusan;default:0"`
	TingkatPartisipasi float64 `gorm:"column:tingkat_partisipasi;type:DECIMAL(5,2);default:0"` // APS
	TingkatAPK         float64 `gorm:"column:tingkat_apk;type:DECIMAL(5,2);default:0"`         // Angka Partisipasi Kasar
	AngkaMelekHuruf    float64 `gorm:"column:angka_melek_huruf;type:DECIMAL(5,2);default:0"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (CapaianPendidikan) TableName() string {
	return "capaian_pendidikan"
}

// ==============================
// 5. Dokumentasi & Laporan
// ==============================
type DokumentasiPendidikan struct {
	ID uint `gorm:"primaryKey;column:id"`

	Judul        string `gorm:"column:judul;type:VARCHAR(150);not null"`
	Keterangan   string `gorm:"column:keterangan;type:TEXT"`
	FilePath     string `gorm:"column:file_path;type:VARCHAR(255)"`     // Dokumen (PDF/Excel)
	FotoKegiatan string `gorm:"column:foto_kegiatan;type:VARCHAR(255)"` // URL foto kegiatan
	Tahun        string `gorm:"column:tahun;type:VARCHAR(4);index"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (DokumentasiPendidikan) TableName() string {
	return "dokumentasi_pendidikan"
}
