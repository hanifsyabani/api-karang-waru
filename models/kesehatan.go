package models

import (
	"time"

	"gorm.io/gorm"
)

type LayananKesehatan struct {
	ID uint `gorm:"primaryKey;column:id"`

	NamaProgram  string `gorm:"column:nama_program;type:VARCHAR(100);not null"`
	Deskripsi    string `gorm:"column:deskripsi;type:TEXT"`
	JenisProgram string `gorm:"column:jenis_program;type:VARCHAR(50)"`
	// contoh: "Imunisasi", "Posyandu", "Posbindu", "Cek Kesehatan"

	FasilitasID *uint              `gorm:"column:fasilitas_id"`
	Fasilitas   FasilitasKesehatan `gorm:"foreignKey:FasilitasID"`

	Jadwal string `gorm:"column:jadwal;type:VARCHAR(200)"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (LayananKesehatan) TableName() string {
	return "layanan_kesehatan"
}

type FasilitasKesehatan struct {
	ID uint `gorm:"primaryKey;column:id"`

	NamaFasilitas string `gorm:"column:nama_fasilitas;type:VARCHAR(100);not null"`
	Jenis         string `gorm:"column:jenis;type:VARCHAR(50);not null"`
	// contoh: "Puskesmas", "Posyandu", "Klinik", "Apotek"

	Alamat          string `gorm:"column:alamat;type:TEXT"`
	PenanggungJawab string `gorm:"column:penanggung_jawab;type:VARCHAR(100)"`

	NoTelepon      string `gorm:"column:no_telepon;type:VARCHAR(20)"`
	JamOperasional string `gorm:"column:jam_operasional;type:VARCHAR(100)"`

	// Metadata
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (FasilitasKesehatan) TableName() string {
	return "fasilitas_kesehatan"
}
