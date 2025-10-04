package responses

import (
	"api-karang-waru/models"
)

type BeritaResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Content   string `json:"content"`
	Excerpt   string `json:"excerpt,omitempty"`
	Writer    string `json:"writer"`
	Image     string `json:"image"`
	Slug      string `json:"slug"`
	Date      string `json:"date"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


func BeritaResponseFromModel(berita *models.Berita) BeritaResponse {
	return BeritaResponse{
		ID:       berita.ID,
		Title:    berita.Title,
		Category: berita.Category,
		Content:  berita.Content,
		Writer:   berita.Writer,
		Image:    berita.Image,
		Status:   berita.Status,
		Slug:     berita.Slug,
		Date:     berita.Date.Format("2006-01-02"),
		CreatedAt: berita.CreatedAt.Format("2006-01-02"),
		UpdatedAt: berita.UpdatedAt.Format("2006-01-02"),
	}
}
