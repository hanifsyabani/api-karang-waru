package responses

import (
	"api-karang-waru/models"
)

type ApbdResponse struct {
	ID    uint `json:"id"`
	Tahun int  `json:"tahun"`

	// Pendapatan Desa
	PendapatanAsliDesa float64 `json:"pendapatan_asli_desa"`
	Transfer           float64 `json:"transfer"`
	PendapatanLain     float64 `json:"pendapatan_lain"`

	// Belanja Desa
	BelanjaPenyelenggaraanPemerintahan float64 `json:"belanja_pemerintahan"`
	BelanjaPembangunan                 float64 `json:"belanja_pembangunan"`
	BelanjaPembinaanKemasyarakatan     float64 `json:"belanja_pembinaan"`
	BelanjaPemberdayaanMasyarakat      float64 `json:"belanja_pemberdayaan"`
	BelanjaTakTerduga                  float64 `json:"belanja_takterduga"`

	// Pembiayaan Desa
	PenerimaanPembiayaan  float64 `json:"penerimaan_pembiayaan"`
	PengeluaranPembiayaan float64 `json:"pengeluaran_pembiayaan"`

	// Total & Status
	TotalPendapatan float64 `json:"total_pendapatan"`
	TotalBelanja    float64 `json:"total_belanja"`
	SurplusDefisit  float64 `json:"surplus_defisit"`
	Status          string  `json:"status"`

	// Metadata
	Keterangan   string `json:"keterangan"`
	FileLampiran string `json:"file_lampiran"`

	// Timestamp
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ApbdResponseFromModel(apbd *models.APBDDesa) ApbdResponse {
	return ApbdResponse{
		ID: apbd.ID,
		Tahun: apbd.Tahun,
		PendapatanAsliDesa: apbd.PendapatanAsliDesa,
		Transfer: apbd.Transfer,
		PendapatanLain: apbd.PendapatanLain,
		BelanjaPenyelenggaraanPemerintahan: apbd.BelanjaPenyelenggaraanPemerintahan,
		BelanjaPembangunan: apbd.BelanjaPembangunan,
		BelanjaPembinaanKemasyarakatan: apbd.BelanjaPembinaanKemasyarakatan,
		BelanjaPemberdayaanMasyarakat: apbd.BelanjaPemberdayaanMasyarakat,
		BelanjaTakTerduga: apbd.BelanjaTakTerduga,
		PenerimaanPembiayaan: apbd.PenerimaanPembiayaan,
		PengeluaranPembiayaan: apbd.PengeluaranPembiayaan,
		TotalPendapatan: apbd.TotalPendapatan,
		TotalBelanja: apbd.TotalBelanja,
		SurplusDefisit: apbd.SurplusDefisit,
		Status: apbd.Status,
		Keterangan: apbd.Keterangan,
		FileLampiran: apbd.FileLampiran,

		CreatedAt: apbd.CreatedAt.Format("2006-01-02"),
		UpdatedAt: apbd.UpdatedAt.Format("2006-01-02"),
	}
}
