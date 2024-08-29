package config

import (
	"os"
)

func LoadConfig() (string, string) {
	secretKey := os.Getenv("SECRET_KEY")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	return secretKey, dbConnectionString
}
