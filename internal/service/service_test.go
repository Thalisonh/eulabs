package service_test

import (
	"testing"

	"github.com/Thalisonh/eulabs.git/internal/database/repository"
	"github.com/Thalisonh/eulabs.git/internal/service"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSave(t *testing.T) {
	fakeProduct := models.Product{
		Name:        "Iphone",
		Price:       4000,
		Description: "Iphone 15",
		Active:      true,
	}

	t.Run("Should return error when fail to save at database", func(t *testing.T) {
		repo := repository.RepositoryMock{
			CreateError: gorm.ErrInvalidTransaction,
		}
		service := service.NewProductService(&repo)

		product, err := service.Save(&fakeProduct)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("Should save at database", func(t *testing.T) {
		fakeId := uuid.New()
		expected := models.Product{
			ID:          fakeId,
			Name:        "Iphone",
			Price:       4000,
			Description: "Iphone 15",
			Active:      true,
		}

		repo := repository.RepositoryMock{
			CreateResult: &expected,
		}

		service := service.NewProductService(&repo)

		product, err := service.Save(&fakeProduct)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, fakeProduct.Name, expected.Name)
	})
}

func TestGetProductById(t *testing.T) {
	t.Run("Should return error when not found", func(t *testing.T) {
		fakeId := uuid.New()

		repo := repository.RepositoryMock{
			GetByIdError: gorm.ErrRecordNotFound,
		}

		service := service.NewProductService(&repo)

		product, err := service.GetProductById(fakeId)
		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("Should a product", func(t *testing.T) {
		fakeId := uuid.New()

		fakeProduct := models.Product{
			ID:          fakeId,
			Name:        "Iphone",
			Price:       4000,
			Description: "Iphone 15",
			Active:      true,
		}

		repo := repository.RepositoryMock{
			GetByIdResult: &fakeProduct,
		}

		service := service.NewProductService(&repo)

		product, err := service.GetProductById(fakeId)
		assert.Nil(t, err)
		assert.NotNil(t, product)
	})
}

func TestGetAllProduct(t *testing.T) {
	t.Run("Should return error when not found", func(t *testing.T) {
		repo := repository.RepositoryMock{
			GetAllError: gorm.ErrRecordNotFound,
		}

		service := service.NewProductService(&repo)

		product, err := service.GetAllProduct()
		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("Should return a list of products", func(t *testing.T) {
		fakeProducts := []models.Product{
			{
				ID:          uuid.New(),
				Name:        "Iphone",
				Price:       4000,
				Description: "Iphone 15",
				Active:      true,
			},
			{
				ID:          uuid.New(),
				Name:        "Iphone",
				Price:       3000,
				Description: "Iphone 14",
				Active:      true,
			},
		}

		repo := repository.RepositoryMock{
			GetAllResult: &fakeProducts,
		}

		service := service.NewProductService(&repo)

		products, err := service.GetAllProduct()
		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Equal(t, 2, len(*products))
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Should return error when not found", func(t *testing.T) {
		fakeId := uuid.New()

		fakeProduct := models.Product{
			Name:        "Iphone",
			Price:       4000,
			Description: "Iphone 15",
			Active:      true,
		}

		repo := repository.RepositoryMock{
			UpdateError: gorm.ErrRecordNotFound,
		}

		service := service.NewProductService(&repo)

		product, err := service.UpdateProduct(fakeId, &fakeProduct)
		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("Should update a product", func(t *testing.T) {
		fakeId := uuid.New()

		fakeProduct := models.Product{
			Name:        "Iphone",
			Price:       4000,
			Description: "Iphone 15",
			Active:      true,
		}

		repo := repository.RepositoryMock{
			UpdateResult: &fakeProduct,
		}

		service := service.NewProductService(&repo)

		products, err := service.UpdateProduct(fakeId, &fakeProduct)
		assert.Nil(t, err)
		assert.NotNil(t, products)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Should return error when fail to delete", func(t *testing.T) {
		fakeId := uuid.New()

		repo := repository.RepositoryMock{
			DeleteError: gorm.ErrInvalidTransaction,
		}

		service := service.NewProductService(&repo)

		err := service.DeleteProduct(fakeId)
		assert.NotNil(t, err)
	})

	t.Run("Should delete a product", func(t *testing.T) {
		fakeId := uuid.New()

		repo := repository.RepositoryMock{
			DeleteError: nil,
		}

		service := service.NewProductService(&repo)

		err := service.DeleteProduct(fakeId)
		assert.Nil(t, err)
	})
}
