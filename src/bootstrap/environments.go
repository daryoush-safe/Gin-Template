package bootstrap

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PRIMARY_DB Database
}

type Database struct {
	DB_HOST string
	DB_NAME string
	DB_PORT string
	DB_USER string
	DB_PASS string
}

func NewEnvironments() *Env {
	godotenv.Load(".env")

	return &Env{
		PRIMARY_DB: Database{
			DB_HOST: os.Getenv("DB_HOST"),
			DB_NAME: os.Getenv("DB_NAME"),
			DB_PORT: os.Getenv("DB_PORT"),
			DB_USER: os.Getenv("DB_USER"),
			DB_PASS: os.Getenv("DB_PASS"),
		},
	}
}
