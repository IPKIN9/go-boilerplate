package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetBaseUrl() (*string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	BaseUrl := os.Getenv("BASE_URL")
	return &BaseUrl, nil
}
