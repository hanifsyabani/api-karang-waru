package requests

type LayananRequest struct {
	ServiceName   string `json:"service_name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Category      string `json:"category" binding:"required"`
	Image         string `json:"image" binding:"required"`
	EstimatedTime string `json:"estimated_time" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Slug          string `json:"slug" binding:"required"`
	Cost          string `json:"cost" binding:"required"`
}
