package server

import (
	"errors"
	"fmt"
	"os"

	"github.com/Thalisonh/eulabs.git/internal/api/routes"
	"github.com/labstack/echo/v4"
)

func Run() error {
	e := echo.New()

	routes.Routes(e)

	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("PORT not found")
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

	return nil
}
