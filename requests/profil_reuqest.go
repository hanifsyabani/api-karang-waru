package requests

type ProfilDesaRequest struct {
	Alamat          string `json:"alamat"`
	Kecamatan       string `json:"kecamatan"`
	Kabupaten       string `json:"kabupaten"`
	Provinsi        string `json:"provinsi"`
	KodePos         string `json:"kode_pos"`
	JumlahPenduduk  string `json:"jumlah_penduduk"`
	JumlahLaki      string `json:"jumlah_laki"`
	JumlahPerempuan string `json:"jumlah_perempuan"`
	JumlahKK        string `json:"jumlah_kk"`

	TahunPembentukan string `json:"tahun_pembentukan"`
	Telepon          string `json:"telepon"`
	Email            string `json:"email"`
}
