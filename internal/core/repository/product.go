package repository

import (
	"errors"
	"fmt"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProducts(offset, limit int) ([]*model.Product, error)
	GetProductById(id int) (*model.Product, error)
	GetProductByName(name string) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product, id *int) error
	DeleteProduct(id *int) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) IProductRepository {
	return &ProductRepository{db: database}
}

func (r *ProductRepository) GetProducts(offset, limit int) ([]*model.Product, error) {
	products := []*model.Product{}

	err := r.db.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get products. Error: %v", err)
	}

	return products, nil
}

func (r *ProductRepository) GetProductById(id int) (*model.Product, error) {
	product := &model.Product{}

	err := r.db.First(product, id).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get product by id. Error: %v", err)

	}

	return product, nil
}

func (r *ProductRepository) GetProductByName(name string) (*model.Product, error) {
	product := &model.Product{}

	err := r.db.Where("name LIKE ?", "%"+name+"%").First(&product).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get product by name. Error: %v", err)
	}

	return product, nil
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	var existingProduct model.Product

	err := r.db.Model(&model.Product{}).Where("name = ?", product.Name).First(&existingProduct).Error
	if err == nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("product already exists. Error: %v", err)
	}

	err = r.db.Create(&product).Error
	if err != nil {
		return fmt.Errorf("fail to create product. Error: %v", err)
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product, id *int) error {
	err := r.db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return fmt.Errorf("fail to update product. Error: %v", err)
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(id *int) error {
	product := &model.Product{
		ID: uint(*id),
	}

	err := r.db.Delete(product).Error
	if err != nil {
		return errors.New("can't delete a product")
	}

	return nil
}
