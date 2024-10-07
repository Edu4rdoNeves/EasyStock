package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID             uint           `gorm:"primarykey;autoIncrement" json:"id"`
	PermissionId   uint64         `gorm:"unique" json:"permission_id"`
	PermissionName string         `gorm:"column:permission_name" json:"permission_name"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (Permission) TableName() string {
	return "Permission"
}
