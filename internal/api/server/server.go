package server

import (
	"github.com/Thalisonh/eulabs.git/internal/api/routes"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
