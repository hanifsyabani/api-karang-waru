package responses

import (
	"api-karang-waru/models"
)

type UmkmResponse struct {
	ID        uint   `json:"id"`
	NamaUsaha string `json:"nama_usaha"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
	Pemilik   string `json:"pemilik"`
	Status    string `json:"status"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


func UmkmResponseFromModel(umkm *models.Umkm) UmkmResponse {
	return UmkmResponse{
		ID:       umkm.ID,
		NamaUsaha: umkm.NamaUsaha,
		Kategori:  umkm.Kategori,
		Deskripsi: umkm.Deskripsi,
		Gambar:    umkm.Gambar,
		Pemilik:   umkm.Pemilik,
		Status:    umkm.Status,
		Slug:      umkm.Slug,
		CreatedAt: umkm.CreatedAt.Format("2006-01-02"),
		UpdatedAt: umkm.UpdatedAt.Format("2006-01-02"),
	}
}
