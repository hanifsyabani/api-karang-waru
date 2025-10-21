package responses

import (
	"api-karang-waru/models"
)

type LayananResponse struct {
	ID            uint   `json:"id"`
	ServiceName   string `json:"service_name"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Image         string `json:"image"`
	EstimatedTime string `json:"estimated_time"`
	Cost          string `json:"cost"`
	Status        string `json:"status"`
	Slug          string `json:"slug"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func LayananResponseFromModel(layanan *models.LayananDesa) LayananResponse {
	return LayananResponse{
		ID:            layanan.ID,
		ServiceName:   layanan.ServiceName,
		Description:   layanan.Description,
		Category:      layanan.Category,
		Image:         layanan.Image,
		EstimatedTime: layanan.EstimatedTime,
		Cost:          layanan.Cost,
		Status:        layanan.Status,
		Slug:          layanan.Slug,
		CreatedAt:     layanan.CreatedAt.Format("2006-01-02"),
		UpdatedAt:     layanan.UpdatedAt.Format("2006-01-02"),
	}
}
