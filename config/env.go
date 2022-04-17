package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var env ENV

// InitDotEnv init params in .env file
func InitDotEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := GetEnvString("PORT")
	database := Database{URI: GetEnvString("DB_URI"), Name: GetEnvString("DB_NAME")}
	jwt := Jwt{SecretKey: GetEnvString("SECRET_KEY")}

	env = ENV{
		AppPort:  appPort,
		Database: database,
		Jwt:      jwt,
	}
}

// GetEnvString ...
func GetEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv return .env data
func GetEnv() *ENV {
	return &env
}
