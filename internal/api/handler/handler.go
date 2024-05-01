package handler

import (
	"errors"
	"net/http"

	"github.com/Thalisonh/eulabs.git/internal/service"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IProductHandler interface {
	AddProduct(c echo.Context) error
	GetProductById(c echo.Context) error
	GetAllProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) IProductHandler {
	return &ProductHandler{service: service}
}

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product
// @Tags product
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /products [post]
func (handler *ProductHandler) AddProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newProduct, err := handler.service.Save(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, newProduct)
}

// GetProductById godoc
// @Summary Get product by id
// @Description Get product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id} [get]
func (handler *ProductHandler) GetProductById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	product, err := handler.service.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "record not found")
		}

		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, product)
}

// GetAllProduct godoc
// @Summary Get all products
// @Description Get all products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {string} string "Internal server error"
// @Router /products [get]
func (handler *ProductHandler) GetAllProduct(c echo.Context) error {
	products, err := handler.service.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id} [put]
func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newProduct, err := handler.service.UpdateProduct(id, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, newProduct)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id} [delete]
func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = handler.service.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.NoContent(http.StatusOK)
}
