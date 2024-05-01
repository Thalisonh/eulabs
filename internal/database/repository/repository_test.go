package repository_test

import (
	"testing"

	"github.com/Thalisonh/eulabs.git/internal/database/repository"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, repository := testDatabase()

	t.Run("Should create a new product", func(t *testing.T) {
		id := uuid.New()

		product := &models.Product{
			ID:          id,
			Name:        "Car",
			Price:       10000.0,
			Description: "Car",
			Active:      true,
		}

		_, err := repository.Create(product)

		assert.Nil(t, err)

		newProduct := &models.Product{}
		err = db.Where("id = ?", id).First(newProduct).Error

		assert.EqualValues(t, product, newProduct)
	})
}

func TestUpdate(t *testing.T) {
	db, repository := testDatabase()

	t.Run("Should update a new product", func(t *testing.T) {
		id := uuid.New()

		oldproduct := &models.Product{
			ID:          id,
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		updated := &models.Product{
			Name:        "New Car",
			Price:       20000.0,
			Description: "New Car",
			Active:      true,
		}

		_, err := repository.Create(oldproduct)

		assert.Nil(t, err)

		_, err = repository.Update(id, updated)

		product := &models.Product{}
		err = db.Where("id = ?", id).First(product).Error

		assert.Equal(t, product.Name, updated.Name)
		assert.Equal(t, product.Active, updated.Active)
		assert.Equal(t, product.Description, updated.Description)
	})
}

func TestDelete(t *testing.T) {
	db, repository := testDatabase()

	t.Run("Should delete a product", func(t *testing.T) {
		id := uuid.New()

		oldProduct := &models.Product{
			ID:          id,
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		_, err := repository.Create(oldProduct)

		assert.Nil(t, err)

		err = repository.Delete(id)

		product := &models.Product{}
		err = db.Where("id = ?", id).First(product).Error

		assert.Equal(t, err, gorm.ErrRecordNotFound)
		assert.ErrorContains(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestGetById(t *testing.T) {
	db, repository := testDatabase()

	t.Run("Should return a product when exist", func(t *testing.T) {
		id := uuid.New()

		newProduct := &models.Product{
			ID:          id,
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		_, err := repository.Create(newProduct)

		assert.Nil(t, err)

		_, err = repository.GetById(id)
		assert.Nil(t, err)

		product := &models.Product{}
		err = db.Where("id = ?", id).First(product).Error

		assert.Equal(t, product, newProduct)
	})

	t.Run("Should return error when product don't exist", func(t *testing.T) {
		id := uuid.New()

		newProduct := &models.Product{
			ID:          id,
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		_, err := repository.Create(newProduct)

		assert.Nil(t, err)

		_, err = repository.GetById(uuid.New())
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	_, repository := testDatabase()

	t.Run("Should return a list of product", func(t *testing.T) {
		product1 := &models.Product{
			ID:          uuid.New(),
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}
		product2 := &models.Product{
			ID:          uuid.New(),
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		_, err := repository.Create(product1)
		assert.Nil(t, err)

		_, err = repository.Create(product2)
		assert.Nil(t, err)

		assert.Nil(t, err)

		products, err := repository.GetAll()
		assert.Nil(t, err)

		assert.Equal(t, 2, len(*products))
	})

	t.Run("Should return error when product don't exist", func(t *testing.T) {
		id := uuid.New()

		newProduct := &models.Product{
			ID:          id,
			Name:        "Old car",
			Price:       10000.0,
			Description: "Old card",
			Active:      true,
		}

		_, err := repository.Create(newProduct)

		assert.Nil(t, err)

		_, err = repository.GetById(uuid.New())
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func testDatabase() (*gorm.DB, repository.IProductRepository) {
	db, err := gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		return nil, nil
	}

	db.AutoMigrate(&models.Product{})

	return db, repository.NewProductRepository(db)
}
