package models

import (
	"time"

	"gorm.io/gorm"
)

// APBDDesa merepresentasikan data anggaran pendapatan dan belanja desa
type APBDDesa struct {
	ID uint `gorm:"primaryKey;column:id"`

	// Tahun Anggaran (contoh: 2024)
	Tahun string `gorm:"column:tahun;type:VARCHAR(4);not null;index"`

	PendapatanAsliDesa float64 `gorm:"column:pendapatan_asli_desa;type:DECIMAL(15,2);default:0"` // PAD
	Transfer           float64 `gorm:"column:transfer;type:DECIMAL(15,2);default:0"`             // Dana transfer dari pusat/provinsi/kabupaten
	PendapatanLain     float64 `gorm:"column:pendapatan_lain;type:DECIMAL(15,2);default:0"`      // Lain-lain pendapatan

	// Belanja Desa
	// =================
	BelanjaPenyelenggaraanPemerintahan float64 `gorm:"column:belanja_pemerintahan;type:DECIMAL(15,2);default:0"`
	BelanjaPembangunan                 float64 `gorm:"column:belanja_pembangunan;type:DECIMAL(15,2);default:0"`
	BelanjaPembinaanKemasyarakatan     float64 `gorm:"column:belanja_pembinaan;type:DECIMAL(15,2);default:0"`
	BelanjaPemberdayaanMasyarakat      float64 `gorm:"column:belanja_pemberdayaan;type:DECIMAL(15,2);default:0"`
	BelanjaTakTerduga                  float64 `gorm:"column:belanja_takterduga;type:DECIMAL(15,2);default:0"`

	// Pembiayaan Desa
	// =================
	PenerimaanPembiayaan  float64 `gorm:"column:penerimaan_pembiayaan;type:DECIMAL(15,2);default:0"`
	PengeluaranPembiayaan float64 `gorm:"column:pengeluaran_pembiayaan;type:DECIMAL(15,2);default:0"`

	// Total & Status
	// =================
	TotalPendapatan float64 `gorm:"column:total_pendapatan;type:DECIMAL(15,2);default:0"`
	TotalBelanja    float64 `gorm:"column:total_belanja;type:DECIMAL(15,2);default:0"`
	SurplusDefisit  float64 `gorm:"column:surplus_defisit;type:DECIMAL(15,2);default:0"`
	Status          string  `gorm:"column:status;type:VARCHAR(50);default:'Published'"` // Published / Draft

	Keterangan   string `gorm:"column:keterangan;type:TEXT"`            // Catatan tambahan (misal: sumber data, keputusan kepala desa)
	FileLampiran string `gorm:"column:file_lampiran;type:VARCHAR(255)"` // File laporan (PDF, Excel, dsb)

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

// TableName menentukan nama tabel di database
func (APBDDesa) TableName() string {
	return "apbd_desa"
}
