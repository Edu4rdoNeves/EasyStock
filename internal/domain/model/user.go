package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"column:Name" json:"name"`
	Email        string         `gorm:"column:Email" json:"email"`
	Password     string         `gorm:"column:Password" json:"password"`
	Permission   Permission     `gorm:"foreignKey:PermissionID" json:"permission"`
	PermissionID uint64         `gorm:"column:permission_id" json:"permission_id"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (Users) TableName() string {
	return "Users"
}
