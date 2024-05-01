package main

import (
	"log"

	"github.com/Thalisonh/eulabs.git/internal/api/server"
	"github.com/joho/godotenv"
)

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
