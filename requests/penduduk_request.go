package requests

// PendudukRequest digunakan untuk menerima input data penduduk dari user (misal via JSON request)
type PendudukRequest struct {
	// Identitas dasar
	NIK          string `json:"nik" binding:"required"`
	NoKK         string `json:"no_kk"`
	NamaLengkap  string `json:"nama_lengkap" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"` // "Laki-laki" / "Perempuan"
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"` // format: "YYYY-MM-DD"

	// Alamat domisili
	Alamat    string `json:"alamat"`
	RT        string `json:"rt"`
	RW        string `json:"rw"`
	Dusun     string `json:"dusun"`
	Desa      string `json:"desa"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`

	// Informasi tambahan
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	Pekerjaan          string `json:"pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	Kewarganegaraan    string `json:"kewarganegaraan"`
	StatusKependudukan string `json:"status_kependudukan"`
	Keterangan          string `json:"keterangan"`
}
