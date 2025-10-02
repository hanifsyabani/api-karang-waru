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
type DemografisRequest struct {
    ProfilDesaID     uint   `json:"profil_desa_id" binding:"required"`
    Balita           string `json:"balita"`
    Anak             string `json:"anak"`
    Dewasa           string `json:"dewasa"`
    Lansia           string `json:"lansia"`
    Pertanian        string `json:"pertanian"`
    Perdagangan      string `json:"perdagangan"`
    Jasa             string `json:"jasa"`
    Industri         string `json:"industri"`
    Sekolah          string `json:"sekolah"`
    Puskesmas        string `json:"puskesmas"`
    Masjid           string `json:"masjid"`
    PasarTradisional string `json:"pasar_tradisional"`
    PosKeamanan      string `json:"pos_keamanan"`
    BalaiDesa        string `json:"balai_desa"`
}

type SejarahRequest struct {
    ProfilDesaID uint   `json:"profil_desa_id" binding:"required"`
    Body         string `json:"body" binding:"required"`
} 


type VisiMisiRequest struct {
    ProfilDesaID uint   `json:"profil_desa_id" binding:"required"`
    Visi        string `json:"visi" binding:"required"`
    Misi        string `json:"misi" binding:"required"`
}
