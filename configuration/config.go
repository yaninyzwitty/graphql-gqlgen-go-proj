package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
	PORT   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	config.DB_URL = getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	config.PORT = getEnv("PORT", "8080")
	return config, nil

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
