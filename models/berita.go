package models

import (
	"time"

	"gorm.io/gorm"
)

type Berita struct {
	ID       uint      `gorm:"primaryKey;column:id"`
	Title    string    `gorm:"column:title;type:VARCHAR(255)"`
	Category string    `gorm:"column:category;type:VARCHAR(100)"`
	Content  string    `gorm:"column:content;type:TEXT"`
	Writer   string    `gorm:"column:writer;type:VARCHAR(100)"`
	Image    string    `gorm:"column:image;type:VARCHAR(255)"`
	Slug     string    `gorm:"column:slug;type:VARCHAR(255);uniqueIndex"`
	Date     time.Time `gorm:"column:date;type:DATE"`
	Status   string    `gorm:"column:status;type:VARCHAR(50)"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

func (Berita) TableName() string {
	return "berita"
}
