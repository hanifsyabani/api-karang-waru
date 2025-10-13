package requests


type UmkmRequest struct {
	NamaUsaha 	 string `json:"nama_usaha" binding:"required"`        // Nama Usaha
	Kategori     string `json:"kategori" binding:"required"`          // Kategori Usaha
	Deskripsi    string `json:"deskripsi" binding:"required"`         // Deskripsi singkat usaha
	Gambar       string `json:"gambar" binding:"required"`            // URL gambar (misal /desa-main.png)
	Pemilik      string `json:"pemilik" binding:"required"`           // Pemilik usaha
	Status       string `json:"status" binding:"required"`            // Status: Terverifikasi / Belum Terverifikasi
	Slug         string `json:"slug" binding:"required"`              // Slug unik
}
