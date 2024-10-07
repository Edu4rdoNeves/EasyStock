package usecases

import (
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
)

type IPermissionUseCases interface {
	GetPermissions(page, limit int) ([]*model.Permission, error)
	GetPermissionById(id string) (*model.Permission, error)
	CreatePermission(Permission *model.Permission) error
	UpdatePermission(Permission *model.Permission, id string) error
	DeletePermission(id string) error
}

type PermissionUseCases struct {
	repository repository.IPermissionRepository
}

func NewPermissionUseCases(repository repository.IPermissionRepository) IPermissionUseCases {
	return &PermissionUseCases{repository}
}

func (u *PermissionUseCases) GetPermissions(page, limit int) ([]*model.Permission, error) {
	offset := (page - 1) * limit

	Permissions, err := u.repository.GetPermissions(offset, limit)
	if err != nil {
		return nil, err
	}

	return Permissions, nil
}

func (u *PermissionUseCases) GetPermissionById(id string) (*model.Permission, error) {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.GetPermissionById(newId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PermissionUseCases) CreatePermission(Permission *model.Permission) error {
	err := u.repository.CreatePermission(Permission)
	if err != nil {
		return err
	}

	return nil
}

func (u *PermissionUseCases) UpdatePermission(Permission *model.Permission, id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = u.repository.UpdatePermission(Permission, &newId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PermissionUseCases) DeletePermission(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = u.repository.DeletePermission(&newId)
	if err != nil {
		return err
	}

	return nil
}
