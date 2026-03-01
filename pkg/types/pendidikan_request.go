package types

//
// ===============================
// 1. Lembaga Pendidikan
// ===============================
//
type LembagaPendidikanRequest struct {
	NamaSekolah       string  `json:"nama_sekolah" binding:"required"`
	JenjangPendidikan string  `json:"jenjang_pendidikan" binding:"required"` // PAUD, SD, SMP, SMA, dll
	Alamat            string  `json:"alamat" binding:"required"`
	JumlahSiswa       int     `json:"jumlah_siswa"`
	JumlahGuru        int     `json:"jumlah_guru"`
	JumlahStaf        int     `json:"jumlah_staf"`
	Kontak            string  `json:"kontak"`
}

//
// ===============================
// 2. Statistik Pendidikan
// ===============================
//
type StatistikPendidikanRequest struct {
	Tahun           string `json:"tahun" binding:"required"` // format YYYY
	TidakSekolah    int    `json:"tidak_sekolah"`
	SD              int    `json:"sd"`
	SMP             int    `json:"smp"`
	SMA             int    `json:"sma"`
	PerguruanTinggi int    `json:"perguruan_tinggi"`
}

//
// ===============================
// 3. Program Pendidikan
// ===============================
//
type ProgramPendidikanRequest struct {
	NamaProgram    string `json:"nama_program" binding:"required"`
	Deskripsi      string `json:"deskripsi"`
	TanggalMulai   string `json:"tanggal_mulai"`   // format: YYYY-MM-DD
	TanggalSelesai string `json:"tanggal_selesai"` // format: YYYY-MM-DD
	Status         string `json:"status"`          // Aktif / Selesai
}

//
// ===============================
// 4. Capaian Pendidikan
// ===============================
//
type CapaianPendidikanRequest struct {
	Tahun              string  `json:"tahun" binding:"required"` // format YYYY
	JumlahLulusan      int     `json:"jumlah_lulusan"`
	TingkatPartisipasi float64 `json:"tingkat_partisipasi"`
	TingkatAPK         float64 `json:"tingkat_apk"`
	AngkaMelekHuruf    float64 `json:"angka_melek_huruf"`
}

//
// ===============================
// 5. Dokumentasi & Laporan Pendidikan
// ===============================
//
type DokumentasiPendidikanRequest struct {
	Judul        string `json:"judul" binding:"required"`
	Keterangan   string `json:"keterangan"`
	FilePath     string `json:"file_path"`     // path dokumen PDF / Excel
	FotoKegiatan string `json:"foto_kegiatan"` // URL gambar
	Tahun        string `json:"tahun"`         // format YYYY
}

