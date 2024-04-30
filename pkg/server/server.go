package server

import (
	"github.com/Thalisonh/eulabs.git/pkg/handler"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.POST("/products", handler.AddProduct)
	e.GET("/products/:id", handler.GetProductById)
	e.GET("/products", handler.GetAllProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
