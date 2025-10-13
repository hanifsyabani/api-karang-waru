package models

import (
	"time"

	"gorm.io/gorm"
)

// Umkm merepresentasikan data UMKM Desa
type Umkm struct {
	ID        uint   `gorm:"primaryKey;column:id"`
	NamaUsaha string `gorm:"column:nama_usaha;type:VARCHAR(150)"` // dari data.title
	Kategori  string `gorm:"column:kategori;type:VARCHAR(100)"`   // dari data.category
	Deskripsi string `gorm:"column:deskripsi;type:TEXT"`          // deskripsi singkat usaha
	Gambar    string `gorm:"column:gambar;type:TEXT"`             // URL gambar (misal /desa-main.png)
	Pemilik   string `gorm:"column:pemilik;type:VARCHAR(100)"`    // contoh: "Kepala Desa"
	Status    string `gorm:"column:status;type:VARCHAR(50)"`      // contoh: "Terverifikasi" atau "Belum Terverifikasi"
	Slug      string `gorm:"column:slug;type:VARCHAR(255);uniqueIndex"`
	

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

// Tentukan nama tabel di database
func (Umkm) TableName() string {
	return "umkm"
}
