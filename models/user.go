package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	Name      string    `gorm:"column:full_name;type:VARCHAR(100);not null"`
	Email     string    `gorm:"column:email;type:VARCHAR(100);not null"`
	Password  string    `gorm:"column:password;type:VARCHAR(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;index"`
}

// define name table
func (User) TableName() string {
	return "users"
}
