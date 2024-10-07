package repository

import (
	"errors"
	"fmt"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUsers(offset, limit int) ([]*model.Users, error)
	GetUserById(id *int) (*model.Users, error)
	CreateUser(user *model.Users) error
	UpdateUser(user *model.Users, id *int) error
	DeleteUser(parseId *int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dataBase *gorm.DB) IUserRepository {
	return &UserRepository{
		db: dataBase,
	}
}

func (r *UserRepository) GetUsers(offset, limit int) ([]*model.Users, error) {
	users := []*model.Users{}

	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get users. Error: %v", err)

	}

	return users, nil
}

func (r *UserRepository) GetUserById(id *int) (*model.Users, error) {
	user := &model.Users{}

	err := r.db.Preload("Permission").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get user by id. Error: %v", err)

	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *model.Users) error {
	var existingUser model.Users

	err := r.db.Model(&model.Users{}).Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user already exists. Error: %v", err)
	}

	err = r.db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("fail to create user. Error: %v", err)
	}
	return nil
}

func (r *UserRepository) UpdateUser(user *model.Users, id *int) error {
	err := r.db.Model(&model.Users{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return fmt.Errorf("fail to update user. Error: %v", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id *int) error {
	user := &model.Users{
		ID: uint64(*id),
	}

	err := r.db.Delete(user).Error
	if err != nil {
		return errors.New("can't delete user")
	}

	return nil
}
