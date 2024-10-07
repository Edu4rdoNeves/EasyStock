package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"gorm.io/gorm"
)

type IPermissionRepository interface {
	GetPermissions(offset, limit int) ([]*model.Permission, error)
	GetPermissionById(id int) (*model.Permission, error)
	CreatePermission(Permission *model.Permission) error
	UpdatePermission(Permission *model.Permission, id *int) error
	DeletePermission(id *int) error
}

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(database *gorm.DB) IPermissionRepository {
	return &PermissionRepository{db: database}
}

func (r *PermissionRepository) GetPermissions(offset, limit int) ([]*model.Permission, error) {
	Permissions := []*model.Permission{}

	err := r.db.Limit(limit).Offset(offset).Find(&Permissions).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get Permissions. Error: %v", err)
	}

	return Permissions, nil
}

func (r *PermissionRepository) GetPermissionById(id int) (*model.Permission, error) {
	Permission := &model.Permission{}

	err := r.db.First(Permission, id).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get Permission by id. Error: %v", err)

	}

	return Permission, nil
}

func (r *PermissionRepository) CreatePermission(Permission *model.Permission) error {
	var existingPermission model.Permission

	lowerPermissionName := strings.ToLower(Permission.PermissionName)

	err := r.db.Model(&model.Permission{}).
		Where("permission_name = ?", lowerPermissionName).
		First(&existingPermission).Error
	if err == nil {
		return fmt.Errorf("permission name already exists. Error: %v", err)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to query permission. Error: %v", err)
	}

	err = r.db.Create(&Permission).Error
	if err != nil {
		return fmt.Errorf("fail to create Permission. Error: %v", err)
	}
	return nil
}

func (r *PermissionRepository) UpdatePermission(Permission *model.Permission, id *int) error {
	err := r.db.Model(&model.Permission{}).Where("id = ?", id).Updates(Permission).Error
	if err != nil {
		return fmt.Errorf("fail to update Permission. Error: %v", err)
	}
	return nil
}

func (r *PermissionRepository) DeletePermission(id *int) error {
	Permission := &model.Permission{
		ID: uint(*id),
	}

	err := r.db.Delete(Permission).Error
	if err != nil {
		return errors.New("can't delete a Permission")
	}

	return nil
}
