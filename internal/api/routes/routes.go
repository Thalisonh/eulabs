package routes

import (
	"github.com/Thalisonh/eulabs.git/internal/api/handler"
	"github.com/Thalisonh/eulabs.git/internal/database"
	"github.com/Thalisonh/eulabs.git/internal/database/repository"
	"github.com/Thalisonh/eulabs.git/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {
	database := database.GetDb()
	repository := repository.NewProductRepository(database)
	service := service.NewProductService(repository)
	handler := handler.NewProductHandler(service)

	e.POST("/products", handler.AddProduct)
	e.GET("/products/:id", handler.GetProductById)
	e.GET("/products", handler.GetAllProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
