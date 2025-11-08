package responses

import "api-karang-waru/models"

// ===============================
// 1. Lembaga Pendidikan
// ===============================
type LembagaPendidikanResponse struct {
	ID                uint    `json:"id"`
	NamaSekolah       string  `json:"nama_sekolah"`
	Alamat            string  `json:"alamat"`
	JenjangPendidikan string  `json:"jenjang_pendidikan"`
	JumlahGuru        int     `json:"jumlah_guru"`
	JumlahSiswa       int     `json:"jumlah_siswa"`
	JumlahStaf        int     `json:"jumlah_staf"`
	Kontak            string  `json:"kontak"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
}

func LembagaPendidikanFromModel(m *models.LembagaPendidikan) LembagaPendidikanResponse {
	return LembagaPendidikanResponse{
		ID: m.ID, NamaSekolah: m.NamaSekolah, Alamat: m.Alamat,
		JenjangPendidikan: m.JenjangPendidikan, JumlahGuru: m.JumlahGuru,
		JumlahSiswa: m.JumlahSiswa, JumlahStaf: m.JumlahStaf,
		Kontak: m.Kontak, Latitude: m.Latitude, Longitude: m.Longitude,
	}
}

// ===============================
// 2. Statistik Pendidikan
// ===============================
type StatistikPendidikanResponse struct {
	ID              uint   `json:"id"`
	Tahun           string `json:"tahun"`
	TidakSekolah    int    `json:"tidak_sekolah"`
	SD              int    `json:"sd"`
	SMP             int    `json:"smp"`
	SMA             int    `json:"sma"`
	PerguruanTinggi int    `json:"perguruan_tinggi"`
}

func StatistikPendidikanFromModel(m *models.StatistikPendidikan) StatistikPendidikanResponse {
	return StatistikPendidikanResponse{
		ID: m.ID, Tahun: m.Tahun, TidakSekolah: m.TidakSekolah,
		SD: m.SD, SMP: m.SMP, SMA: m.SMA, PerguruanTinggi: m.PerguruanTinggi,
	}
}

// ===============================
// 3. Program Pendidikan
// ===============================
type ProgramPendidikanResponse struct {
	ID             uint   `json:"id"`
	NamaProgram    string `json:"nama_program"`
	Deskripsi      string `json:"deskripsi"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	Status         string `json:"status"`
}

func ProgramPendidikanFromModel(m *models.ProgramPendidikan) ProgramPendidikanResponse {
	return ProgramPendidikanResponse{
		ID: m.ID, NamaProgram: m.NamaProgram, Deskripsi: m.Deskripsi,
		TanggalMulai:   m.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: m.TanggalSelesai.Format("2006-01-02"),
		Status:         m.Status,
	}
}

// ===============================
// 4. Capaian Pendidikan
// ===============================
type CapaianPendidikanResponse struct {
	ID                 uint    `json:"id"`
	JumlahLulusan      int     `json:"jumlah_lulusan"`
	TingkatPartisipasi float64 `json:"tingkat_partisipasi"`
	TingkatAPK         float64 `json:"tingkat_apk"`
	AngkaMelekHuruf    float64 `json:"angka_melek_huruf"`
	Tahun              string  `json:"tahun"`
}

func CapaianPendidikanFromModel(m *models.CapaianPendidikan) CapaianPendidikanResponse {
	return CapaianPendidikanResponse{
		ID: m.ID, JumlahLulusan: m.JumlahLulusan, TingkatPartisipasi: m.TingkatPartisipasi,
		TingkatAPK: m.TingkatAPK, AngkaMelekHuruf: m.AngkaMelekHuruf, Tahun: m.Tahun,
	}
}

// ===============================
// 5. Dokumentasi Pendidikan
// ===============================
type DokumentasiPendidikanResponse struct {
	ID           uint   `json:"id"`
	Judul        string `json:"judul"`
	Keterangan   string `json:"keterangan"`
	FilePath     string `json:"file_path"`
	FotoKegiatan string `json:"foto_kegiatan"`
	Tahun        string    `json:"tahun"`
}

func DokumentasiPendidikanFromModel(m *models.DokumentasiPendidikan) DokumentasiPendidikanResponse {
	return DokumentasiPendidikanResponse{
		ID: m.ID, Judul: m.Judul, Keterangan: m.Keterangan,
		FilePath: m.FilePath, FotoKegiatan: m.FotoKegiatan, Tahun: m.Tahun,
	}
}
