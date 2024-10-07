package usecases

import (
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
)

type IProductUseCases interface {
	GetProducts(page, limit int) ([]*model.Product, error)
	GetProductById(id string) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product, id string) error
	DeleteProduct(id string) error
}

type ProductUseCases struct {
	repository repository.IProductRepository
}

func NewProductUseCases(repository repository.IProductRepository) IProductUseCases {
	return &ProductUseCases{repository}
}

func (u *ProductUseCases) GetProducts(page, limit int) ([]*model.Product, error) {
	offset := (page - 1) * limit

	products, err := u.repository.GetProducts(offset, limit)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUseCases) GetProductById(id string) (*model.Product, error) {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.GetProductById(newId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *ProductUseCases) CreateProduct(product *model.Product) error {
	err := u.repository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUseCases) UpdateProduct(product *model.Product, id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = u.repository.UpdateProduct(product, &newId)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUseCases) DeleteProduct(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = u.repository.DeleteProduct(&newId)
	if err != nil {
		return err
	}

	return nil
}
