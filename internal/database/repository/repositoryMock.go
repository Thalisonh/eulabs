package repository

import (
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
)

type RepositoryMock struct {
	CreateResult  *models.Product
	CreateError   error
	UpdateResult  *models.Product
	UpdateError   error
	GetByIdResult *models.Product
	GetByIdError  error
	GetAllResult  *[]models.Product
	GetAllError   error
	DeleteError   error
}

func (mock *RepositoryMock) Create(product *models.Product) (*models.Product, error) {
	return mock.CreateResult, mock.CreateError
}

func (mock *RepositoryMock) Update(id uuid.UUID, product *models.Product) (*models.Product, error) {
	return mock.UpdateResult, mock.UpdateError
}

func (mock *RepositoryMock) GetById(id uuid.UUID) (*models.Product, error) {
	return mock.GetByIdResult, mock.GetByIdError
}

func (mock *RepositoryMock) GetAll() (*[]models.Product, error) {
	return mock.GetAllResult, mock.GetAllError
}

func (mock *RepositoryMock) Delete(id uuid.UUID) error {
	return mock.DeleteError
}
