package requests


type BeritaRequest struct {
    Title    string `json:"title" binding:"required"`      // judul berita
    Category string `json:"category" binding:"required"`   // kategori: Pengumuman, Agenda, dll
    Content  string `json:"content" binding:"required"`    // isi berita
    Writer   string `json:"writer"`                        // penulis
    Image    string `json:"image"`                         // url gambar
    Slug     string `json:"slug"`                          // slug unik
    Date     string `json:"date"`                          // tanggal berita (YYYY-MM-DD)
    Status   string `json:"status" binding:"omitempty"`    // draft / publish
}
