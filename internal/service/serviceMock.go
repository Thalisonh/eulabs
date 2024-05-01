package service

import (
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
)

type ServiceMock struct {
	SaveResult           *models.Product
	SaveError            error
	GetProductByIdResult *models.Product
	GetProductByIdError  error
	GetAllProductResult  *[]models.Product
	GetAllProductError   error
	UpdateProductResult  *models.Product
	UpdateProductError   error
	DeleteProductError   error
}

func (mock *ServiceMock) Save(product *models.Product) (*models.Product, error) {
	return mock.SaveResult, mock.SaveError
}

func (mock *ServiceMock) GetProductById(id uuid.UUID) (*models.Product, error) {
	return mock.GetProductByIdResult, mock.GetProductByIdError
}

func (mock *ServiceMock) GetAllProduct() (*[]models.Product, error) {
	return mock.GetAllProductResult, mock.GetAllProductError
}

func (mock *ServiceMock) UpdateProduct(id uuid.UUID, product *models.Product) (*models.Product, error) {
	return mock.UpdateProductResult, mock.UpdateProductError
}

func (mock *ServiceMock) DeleteProduct(id uuid.UUID) error {
	return mock.DeleteProductError
}
