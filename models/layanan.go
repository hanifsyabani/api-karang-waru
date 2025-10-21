package models

import (
	"time"

	"gorm.io/gorm"
)

type LayananDesa struct {
	ID            uint   `gorm:"primaryKey;column:id"`
	ServiceName   string `gorm:"column:service_name;type:VARCHAR(255);not null"` // Service title, e.g., "Business Permit Application"
	Description   string `gorm:"column:description;type:TEXT"`                   // Detailed explanation, requirements, and workflow
	Category      string `gorm:"column:category;type:VARCHAR(100)"`              // e.g., "Population", "Administration", "Social"
	Image         string `gorm:"column:image;type:VARCHAR(255)"`                 // Path to an icon or image for the service
	Slug          string `gorm:"column:slug;type:VARCHAR(255);uniqueIndex"`      // For clean URLs (e.g., /services/business-permit)
	Status        string `gorm:"column:status;type:VARCHAR(50)"`                 // e.g., "Published", "Draft" (controls visibility on the landing page)
	
	EstimatedTime string `gorm:"column:estimated_time;type:VARCHAR(100)"`      
	Cost          string `gorm:"column:cost;type:VARCHAR(100)"`                  

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (LayananDesa) TableName() string {
	return "village_services"
}
