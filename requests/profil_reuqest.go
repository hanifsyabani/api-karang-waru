package requests

type ProfilDesaRequest struct {
	Alamat           string  `json:"alamat"`
	Kecamatan        string  `json:"kecamatan"`
	Kabupaten        string  `json:"kabupaten"`
	Provinsi         string  `json:"provinsi"`
	KodePos          string  `json:"kode_pos"`
	JumlahPenduduk   int     `json:"jumlah_penduduk"`
	JumlahLaki       int     `json:"jumlah_laki"`
	JumlahPerempuan  int     `json:"jumlah_perempuan"`
	JumlahKK         int     `json:"jumlah_kk"`
	LuasWilayahKm2   float64 `json:"luas_wilayah_km2"`
	LuasWilayahHa    int     `json:"luas_wilayah_ha"`
	TahunPembentukan int     `json:"tahun_pembentukan"`
	Telepon          string  `json:"telepon"`
	Email            string  `json:"email"`
	JamPelayanan     string  `json:"jam_pelayanan"`
}