package repository

import (
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	Update(id int, product *models.Product) (*models.Product, error)
	GetById(id int) (*models.Product, error)
	GetAll() (*models.Product, error)
	Delete() error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	return nil, nil
}

func (repository *ProductRepository) Update(id int, product *models.Product) (*models.Product, error) {
	return nil, nil
}

func (repository *ProductRepository) GetById(id int) (*models.Product, error) { return nil, nil }

func (repository *ProductRepository) GetAll() (*models.Product, error) { return nil, nil }

func (repository *ProductRepository) Delete() error { return nil }
