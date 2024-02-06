package main

import (
	"log"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/apiserver"
)

var (
	configPath       string
	configPathDocker string
)

func main() {
	config := apiserver.ConfigLoad()

	err := apiserver.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
