package repository

import (
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	Update(id uuid.UUID, product *models.Product) (*models.Product, error)
	GetById(id uuid.UUID) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	Delete(id uuid.UUID) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	err := repository.db.Create(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) Update(id uuid.UUID, product *models.Product) (*models.Product, error) {
	err := repository.db.Where("id = ?", id).Updates(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) GetById(id uuid.UUID) (*models.Product, error) {
	product := &models.Product{}
	err := repository.db.Where("id = ?", id).First(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) GetAll() (*[]models.Product, error) {
	products := &[]models.Product{}
	err := repository.db.Find(products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *ProductRepository) Delete(id uuid.UUID) error {
	err := repository.db.Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}
