package database

import (
	"log"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {

	err := db.AutoMigrate(
		&model.Users{},
		&model.Product{},
		&model.Permission{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
