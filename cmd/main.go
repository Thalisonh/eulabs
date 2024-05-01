package main

import (
	"log"

	_ "github.com/Thalisonh/eulabs.git/docs"
	"github.com/Thalisonh/eulabs.git/internal/api/server"
	"github.com/joho/godotenv"
)

// @title Swagger Eulabs API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	errDotEnv := godotenv.Load()

	if errDotEnv != nil {
		log.Fatal("Error loading .env files")
	}

	err := server.Run()
	if err != nil {
		panic(err)
	}

}
