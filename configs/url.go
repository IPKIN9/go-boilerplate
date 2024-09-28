package consts

import "os"

func GetBaseUrl() (*string, error) {
	BaseUrl := os.Getenv("GO_BASE_URL")
	return &BaseUrl, nil
}
