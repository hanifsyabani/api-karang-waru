package domain

import (
	"time"

	"gorm.io/gorm"
)

// Penduduk merepresentasikan data penduduk desa
type Penduduk struct {
	ID uint `gorm:"primaryKey;column:id"`

	// Identitas dasar
	NIK          string    `gorm:"column:nik;type:VARCHAR(20);unique;not null"` // Nomor Induk Kependudukan
	NoKK         string    `gorm:"column:no_kk;type:VARCHAR(20)"`               // Nomor Kartu Keluarga
	NamaLengkap  string    `gorm:"column:nama_lengkap;type:VARCHAR(100);not null"`
	JenisKelamin string    `gorm:"column:jenis_kelamin;type:VARCHAR(20)"` // Contoh: "Laki-laki" / "Perempuan"
	TempatLahir  string    `gorm:"column:tempat_lahir;type:VARCHAR(50)"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;type:DATE"`

	// Alamat domisili
	Alamat    string `gorm:"column:alamat;type:TEXT"`
	RT        string `gorm:"column:rt;type:VARCHAR(3)"`
	RW        string `gorm:"column:rw;type:VARCHAR(3)"`
	Dusun     string `gorm:"column:dusun;type:VARCHAR(50)"`
	Desa      string `gorm:"column:desa;type:VARCHAR(100)"`
	Kecamatan string `gorm:"column:kecamatan;type:VARCHAR(100)"`
	Kabupaten string `gorm:"column:kabupaten;type:VARCHAR(100)"`
	Provinsi  string `gorm:"column:provinsi;type:VARCHAR(100)"`

	// Informasi tambahan
	Agama              string `gorm:"column:agama;type:VARCHAR(50)"`
	StatusPerkawinan   string `gorm:"column:status_perkawinan;type:VARCHAR(50)"`
	Pekerjaan          string `gorm:"column:pekerjaan;type:VARCHAR(100)"`
	PendidikanTerakhir string `gorm:"column:pendidikan_terakhir;type:VARCHAR(100)"`
	Kewarganegaraan    string `gorm:"column:kewarganegaraan;type:VARCHAR(50);default:'WNI'"`

	// Status kependudukan
	StatusKependudukan string `gorm:"column:status_kependudukan;type:VARCHAR(50);default:'Tetap'"`
	Keterangan          string `gorm:"column:keterangan;type:TEXT"`

	// Metadata
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

// TableName menentukan nama tabel database
func (Penduduk) TableName() string {
	return "penduduk"
}

