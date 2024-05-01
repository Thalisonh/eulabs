package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Thalisonh/eulabs.git/internal/api/handler"
	"github.com/Thalisonh/eulabs.git/internal/service"
	"github.com/Thalisonh/eulabs.git/pkg/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddProduct(t *testing.T) {
	t.Run("Should return error when payload is invalid", func(t *testing.T) {
		service := service.ServiceMock{}

		handler := handler.NewProductHandler(&service)

		expected := map[string]string{
			"Name":        "Iphone",
			"Description": "Iphone 15",
			"Price":       "4000",
			"Active":      "true",
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.AddProduct(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Should return error when fail to call service", func(t *testing.T) {
		service := service.ServiceMock{
			SaveError: errors.New("internal server error"),
		}

		handler := handler.NewProductHandler(&service)

		expected := models.Product{
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.AddProduct(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Should create a product", func(t *testing.T) {
		service := service.ServiceMock{
			SaveResult: &models.Product{
				ID:          uuid.New(),
				Name:        "Iphone",
				Description: "Iphone 15",
				Price:       4000,
				Active:      true,
			},
		}

		handler := handler.NewProductHandler(&service)

		expected := models.Product{
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.AddProduct(c)) {
			product := &models.Product{}
			json.Unmarshal(rec.Body.Bytes(), product)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected.Name, product.Name)
			assert.Equal(t, expected.Description, product.Description)
			assert.Equal(t, expected.Price, product.Price)
			assert.Equal(t, expected.Active, product.Active)
		}
	})
}

func TestGetProductById(t *testing.T) {
	t.Run("Should return error when id is not a uuid valid", func(t *testing.T) {
		service := service.ServiceMock{}
		handler := handler.NewProductHandler(&service)

		fakeId := "1"

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId)

		if assert.NoError(t, handler.GetProductById(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Should return error when not found product by id", func(t *testing.T) {
		service := service.ServiceMock{
			GetProductByIdError: gorm.ErrRecordNotFound,
		}
		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New().String()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId)

		if assert.NoError(t, handler.GetProductById(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("Should return error fail to call service", func(t *testing.T) {
		service := service.ServiceMock{
			GetProductByIdError: errors.New("internal server error"),
		}
		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New().String()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId)

		if assert.NoError(t, handler.GetProductById(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Should return a product", func(t *testing.T) {
		fakeId := uuid.New()
		expected := &models.Product{
			ID:          fakeId,
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		service := service.ServiceMock{
			GetProductByIdResult: expected,
		}
		handler := handler.NewProductHandler(&service)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.GetProductById(c)) {
			product := &models.Product{}
			json.Unmarshal(rec.Body.Bytes(), product)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, product, expected)
		}
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Should return error fail to call service", func(t *testing.T) {
		service := service.ServiceMock{
			GetAllProductError: gorm.ErrNotImplemented,
		}
		handler := handler.NewProductHandler(&service)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.GetAllProduct(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Should return a list of products", func(t *testing.T) {
		expected := &[]models.Product{{
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
			}}

		service := service.ServiceMock{
			GetAllProductResult: expected,
		}
		handler := handler.NewProductHandler(&service)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.GetAllProduct(c)) {
			products := &[]models.Product{}
			json.Unmarshal(rec.Body.Bytes(), products)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, products, expected)
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Should return error when id is not a valid uuid", func(t *testing.T) {
		service := service.ServiceMock{}

		handler := handler.NewProductHandler(&service)

		expected := models.Product{
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		b, _ := json.Marshal(expected)

		fakeId := "1"

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/:id", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId)

		if assert.NoError(t, handler.UpdateProduct(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Should return error when payload is invalid", func(t *testing.T) {
		service := service.ServiceMock{}

		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New()

		expected := map[string]string{
			"Name":        "Iphone",
			"Description": "Iphone 15",
			"Price":       "4000",
			"Active":      "true",
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/:id", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.UpdateProduct(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Should return error when fail to call service", func(t *testing.T) {
		service := service.ServiceMock{
			UpdateProductError: errors.New("internal server error"),
		}

		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New()

		expected := models.Product{
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/:id", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.UpdateProduct(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Should update a product", func(t *testing.T) {
		service := service.ServiceMock{
			UpdateProductResult: &models.Product{
				ID:          uuid.New(),
				Name:        "Iphone",
				Description: "Iphone 15",
				Price:       4000,
				Active:      true,
			},
		}

		handler := handler.NewProductHandler(&service)
		fakeId := uuid.New()

		expected := models.Product{
			ID:          fakeId,
			Name:        "Iphone",
			Description: "Iphone 15",
			Price:       4000,
			Active:      true,
		}

		b, _ := json.Marshal(expected)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/products/:id", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.UpdateProduct(c)) {
			product := &models.Product{}
			json.Unmarshal(rec.Body.Bytes(), product)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected.Name, product.Name)
			assert.Equal(t, expected.Description, product.Description)
			assert.Equal(t, expected.Price, product.Price)
			assert.Equal(t, expected.Active, product.Active)
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Should return error when id is not a valid uuid", func(t *testing.T) {
		service := service.ServiceMock{}

		handler := handler.NewProductHandler(&service)

		fakeId := "1"

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId)

		if assert.NoError(t, handler.DeleteProduct(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Should return error when fail to call service", func(t *testing.T) {
		service := service.ServiceMock{
			DeleteProductError: errors.New("internal server error"),
		}

		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New()

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.DeleteProduct(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Should delete a product", func(t *testing.T) {
		service := service.ServiceMock{
			DeleteProductError: nil,
		}

		handler := handler.NewProductHandler(&service)

		fakeId := uuid.New()

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fakeId.String())

		if assert.NoError(t, handler.DeleteProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
