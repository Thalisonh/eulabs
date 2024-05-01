package service

import (
	"github.com/Thalisonh/eulabs.git/internal/database/repository"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
)

type IProductService interface {
	Save(product *models.Product) (*models.Product, error)
	GetProductById(id uuid.UUID) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
	UpdateProduct(id uuid.UUID, product *models.Product) (*models.Product, error)
	DeleteProduct(id uuid.UUID) error
}

type ProductService struct {
	repository repository.IProductRepository
}

func NewProductService(repository repository.IProductRepository) IProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) Save(product *models.Product) (*models.Product, error) {
	product.ID = uuid.New()

	newProduct, err := service.repository.Create(product)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (service *ProductService) GetProductById(id uuid.UUID) (*models.Product, error) {
	product, err := service.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) GetAllProduct() (*[]models.Product, error) {
	products, err := service.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductService) UpdateProduct(id uuid.UUID, product *models.Product) (*models.Product, error) {
	_, err := service.repository.Update(id, product)
	if err != nil {
		return nil, err
	}

	newProduct, err := service.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (service *ProductService) DeleteProduct(id uuid.UUID) error {
	err := service.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
