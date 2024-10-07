package usecases

import (
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
)

type IUserUseCases interface {
	GetUsers(page, limit int) ([]*model.Users, error)
	GetUserById(id string) (*model.Users, error)
	CreateUser(user *model.Users) error
	UpdateUser(user *model.Users, id string) error
	DeleteUser(id string) error
}
type UserUseCases struct {
	repository repository.IUserRepository
}

func NewUserUseCases(repository repository.IUserRepository) IUserUseCases {
	return &UserUseCases{repository}
}

func (b *UserUseCases) GetUsers(page, limit int) ([]*model.Users, error) {
	offset := (page - 1) * limit

	userResponse, err := b.repository.GetUsers(offset, limit)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (b *UserUseCases) GetUserById(id string) (*model.Users, error) {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user, err := b.repository.GetUserById(&newId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (b *UserUseCases) CreateUser(user *model.Users) error {
	if user.PermissionID == 0 {
		user.PermissionID = 2
	}

	err := b.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (b *UserUseCases) UpdateUser(user *model.Users, id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = b.repository.UpdateUser(user, &newId)
	if err != nil {
		return err
	}

	return nil
}

func (b *UserUseCases) DeleteUser(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = b.repository.DeleteUser(&newId)
	if err != nil {
		return err
	}

	return nil
}
