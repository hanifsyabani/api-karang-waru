package requests

// ApbdRequest digunakan untuk menerima input data APBD Desa dari user (misal via JSON request)
type ApbdRequest struct {
	Tahun string `json:"tahun" binding:"required"`

	// Pendapatan Desa
	PendapatanAsliDesa float64 `json:"pendapatan_asli_desa"`
	Transfer            float64 `json:"transfer"`
	PendapatanLain      float64 `json:"pendapatan_lain"`

	// Belanja Desa
	BelanjaPenyelenggaraanPemerintahan float64 `json:"belanja_pemerintahan"`
	BelanjaPembangunan                 float64 `json:"belanja_pembangunan"`
	BelanjaPembinaanKemasyarakatan     float64 `json:"belanja_pembinaan"`
	BelanjaPemberdayaanMasyarakat      float64 `json:"belanja_pemberdayaan"`
	BelanjaTakTerduga                  float64 `json:"belanja_takterduga"`

	// Pembiayaan Desa
	PenerimaanPembiayaan  float64 `json:"penerimaan_pembiayaan"`
	PengeluaranPembiayaan float64 `json:"pengeluaran_pembiayaan"`

	Status          string  `json:"status"`

	// Metadata
	Keterangan   string `json:"keterangan"`
	FileLampiran string `json:"file_lampiran"`
}
