package models

import (
	"time"

	"gorm.io/gorm"
)

type ProfilDesa struct {
	ID               uint   `gorm:"primaryKey;column:id"`
	Alamat           string `gorm:"column:alamat;type:TEXT"`
	Kecamatan        string `gorm:"column:kecamatan;type:VARCHAR(100)"`
	Kabupaten        string `gorm:"column:kabupaten;type:VARCHAR(100)"`
	Provinsi         string `gorm:"column:provinsi;type:VARCHAR(100)"`
	KodePos          string `gorm:"column:kode_pos;type:VARCHAR(10)"`
	JumlahPenduduk   string `gorm:"column:jumlah_penduduk;type:VARCHAR(10)"`
	JumlahLaki       string `gorm:"column:jumlah_laki;type:VARCHAR(10)"`
	JumlahPerempuan  string `gorm:"column:jumlah_perempuan;type:VARCHAR(10)"`
	JumlahKK         string `gorm:"column:jumlah_kk;type:VARCHAR(10)"`
	TahunPembentukan string `gorm:"column:tahun_pembentukan;type:VARCHAR(4)"`
	Telepon          string `gorm:"column:telepon;type:VARCHAR(20)"`
	Email            string `gorm:"column:email;type:VARCHAR(150)"`

	Demografis []Demografis `gorm:"foreignKey:ProfilDesaID"`
	Sejarah    []Sejarah    `gorm:"foreignKey:ProfilDesaID"`
	VisiMisi   []VisiMisi   `gorm:"foreignKey:ProfilDesaID"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (ProfilDesa) TableName() string {
	return "profil_desa"
}

type Demografis struct {
	ID               uint   `gorm:"primaryKey;column:id"`
	ProfilDesaID     uint   `gorm:"column:profil_desa_id;type:BIGINT;index"`
	Balita           string `gorm:"column:balita;type:VARCHAR(50)"`
	Anak             string `gorm:"column:anak;type:VARCHAR(50)"`
	Dewasa           string `gorm:"column:dewasa;type:VARCHAR(50)"`
	Lansia           string `gorm:"column:lansia;type:VARCHAR(50)"`
	Pertanian        string `gorm:"column:pertanian;type:VARCHAR(50)"`
	Perdagangan      string `gorm:"column:perdagangan;type:VARCHAR(50)"`
	Jasa             string `gorm:"column:jasa;type:VARCHAR(50)"`
	Industri         string `gorm:"column:industri;type:VARCHAR(50)"`
	Sekolah          string `gorm:"column:sekolah;type:VARCHAR(50)"`
	Puskesmas        string `gorm:"column:puskesmas;type:VARCHAR(50)"`
	Masjid           string `gorm:"column:masjid;type:VARCHAR(50)"`
	PasarTradisional string `gorm:"column:pasar_tradisional;type:VARCHAR(50)"`
	PosKeamanan      string `gorm:"column:pos_keamanan;type:VARCHAR(50)"`
	BalaiDesa        string `gorm:"column:balai_desa;type:VARCHAR(50)"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (Demografis) TableName() string {
	return "demografis"
}

type Sejarah struct {
	ID           uint   `gorm:"primaryKey;column:id"`
	ProfilDesaID uint   `gorm:"column:profil_desa_id;type:BIGINT;index"`
	Body         string `gorm:"column:body;type:TEXT"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (Sejarah) TableName() string {
	return "sejarah"
}

type VisiMisi struct {
	ID           uint   `gorm:"primaryKey;column:id"`
	ProfilDesaID uint   `gorm:"column:profil_desa_id;type:BIGINT;index"`
	Visi         string `gorm:"column:visi;type:TEXT"`
	Misi         string `gorm:"column:misi;type:TEXT"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (VisiMisi) TableName() string {
	return "visi_misi"
}
