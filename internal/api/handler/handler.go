package handler

import (
	"net/http"
	"strconv"

	"github.com/Thalisonh/eulabs.git/internal/service"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/labstack/echo/v4"
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

func (handler *ProductHandler) AddProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newProduct, err := handler.service.Save(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, newProduct)
}

func (handler *ProductHandler) GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	product, err := handler.service.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) GetAllProduct(c echo.Context) error {
	products, err := handler.service.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newProduct, err := handler.service.UpdateProduct(id, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, newProduct)
}

func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = handler.service.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.NoContent(http.StatusOK)
}
