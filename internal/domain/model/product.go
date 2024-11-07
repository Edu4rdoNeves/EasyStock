package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey;autoIncrement" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description"`
	Price       string         `gorm:"column:price" json:"price"`
	Quantity    int64          `gorm:"column:quantity" json:"quantity"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (Product) TableName() string {
	return "Products"
}
