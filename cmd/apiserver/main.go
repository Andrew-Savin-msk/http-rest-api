package main

import (
	"fmt"
	"log"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/apiserver"
)

var (
	configPath       string
	configPathDocker string
)

func main() {
	fmt.Println("Work Starts")

	config := apiserver.ConfigLoad()
	fmt.Println("Config loaded")

	err := apiserver.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
