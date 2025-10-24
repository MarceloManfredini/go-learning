package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables from .env file
	err := godotenv.Load("C:\\Marcelo\\GO\\Exercicios\\Crud" + `\.env`)

	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
}
