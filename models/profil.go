package models

import (
	"time"

	"gorm.io/gorm"
)

type ProfilDesa struct {
	ID               uint              `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	Alamat           string            `gorm:"column:alamat;type:TEXT"`
	Kecamatan        string            `gorm:"column:kecamatan;type:VARCHAR(100)"`
	Kabupaten        string            `gorm:"column:kabupaten;type:VARCHAR(100)"`
	Provinsi         string            `gorm:"column:provinsi;type:VARCHAR(100)"`
	KodePos          string            `gorm:"column:kode_pos;type:VARCHAR(10)"`
	JumlahPenduduk   int               `gorm:"column:jumlah_penduduk;type:INT"`
	JumlahLaki       int               `gorm:"column:jumlah_laki;type:INT"`
	JumlahPerempuan  int               `gorm:"column:jumlah_perempuan;type:INT"`
	JumlahKK         int               `gorm:"column:jumlah_kk;type:INT"`
	LuasWilayahKm2   float64           `gorm:"column:luas_wilayah_km2;type:DECIMAL(6,2)"`
	LuasWilayahHa    int               `gorm:"column:luas_wilayah_ha;type:INT"`
	TahunPembentukan int               `gorm:"column:tahun_pembentukan;type:YEAR"`
	Telepon          string            `gorm:"column:telepon;type:VARCHAR(20)"`
	Email            string            `gorm:"column:email;type:VARCHAR(150)"`
	JamPelayanan     string            `gorm:"column:jam_pelayanan;type:JSON"`

	// Relasi
	BatasWilayah      []BatasWilayah      `gorm:"foreignKey:ProfilDesaID"`
	KomposisiPenduduk []KomposisiPenduduk `gorm:"foreignKey:ProfilDesaID"`
	MataPencaharian   []MataPencaharian   `gorm:"foreignKey:ProfilDesaID"`
	FasilitasUmum     []FasilitasUmum     `gorm:"foreignKey:ProfilDesaID"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (ProfilDesa) TableName() string {
	return "profil_desa"
}

type BatasWilayah struct {
	ID           uint   `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	ProfilDesaID uint   `gorm:"column:profil_desa_id;type:BIGINT UNSIGNED;index"`
	Arah         string `gorm:"column:arah;type:ENUM('Utara','Selatan','Timur','Barat')"`
	DesaBatas    string `gorm:"column:desa_batas;type:VARCHAR(100)"`
}

func (BatasWilayah) TableName() string {
	return "batas_wilayah"
}

type KomposisiPenduduk struct {
	ID           uint   `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	ProfilDesaID uint   `gorm:"column:profil_desa_id;type:BIGINT UNSIGNED;index"`
	Kategori     string `gorm:"column:kategori;type:ENUM('Balita','Anak-anak','Dewasa','Lansia')"`
	Jumlah       int    `gorm:"column:jumlah;type:INT"`
}

func (KomposisiPenduduk) TableName() string {
	return "komposisi_penduduk"
}

type MataPencaharian struct {
	ID           uint    `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	ProfilDesaID uint    `gorm:"column:profil_desa_id;type:BIGINT UNSIGNED;index"`
	Jenis        string  `gorm:"column:jenis;type:ENUM('Pertanian','Perdagangan','Jasa','Industri Kecil')"`
	Persentase   float64 `gorm:"column:persentase;type:DECIMAL(5,2)"`
}

func (MataPencaharian) TableName() string {
	return "mata_pencaharian"
}

type FasilitasUmum struct {
	ID           uint   `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	ProfilDesaID uint   `gorm:"column:profil_desa_id;type:BIGINT UNSIGNED;index"`
	Jenis        string `gorm:"column:jenis;type:ENUM('Sekolah','Puskesmas','Masjid','Pasar Tradisional','Pos Keamanan','Balai Desa')"`
	Jumlah       int    `gorm:"column:jumlah;type:INT"`
}

func (FasilitasUmum) TableName() string {
	return "fasilitas_umum"
}
