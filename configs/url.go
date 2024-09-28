package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}
}

func GetBaseUrl() *string {
	BaseUrl := os.Getenv("BASE_URL")
	if len(BaseUrl) < 1 {
		return nil
	}
	return &BaseUrl
}
