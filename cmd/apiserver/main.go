package main

import (
	"flag"
	"log"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath       string
	configPathDocker string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&configPathDocker, "config-path-docker", "configs/apiserverDocker.toml", "path to config file for docker DB")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPathDocker, config)
	if err != nil {
		log.Fatal(err)
	}

	err = apiserver.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
