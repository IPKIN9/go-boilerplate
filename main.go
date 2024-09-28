package main

import (
	"fmt"

	"github.com/IPKIN9/go-boilerplate/configs"
)

func main() {
	baseUrl, err := configs.GetBaseUrl()

	fmt.Println("Try to get error, \n", err)
	fmt.Println("Try to get config, \n", *baseUrl)
}
