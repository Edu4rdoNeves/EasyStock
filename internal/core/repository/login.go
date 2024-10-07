package repository

import (
	"errors"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"gorm.io/gorm"
)

type ILoginRepository interface {
	Login(login *model.Login, user *model.Users) error
}

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(database *gorm.DB) ILoginRepository {
	return &LoginRepository{
		db: database,
	}
}

func (r *LoginRepository) Login(login *model.Login, user *model.Users) error {
	dbError := r.db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		return errors.New("error: Invalid Email")
	}

	return nil
}
