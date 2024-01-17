package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	fmt.Println(*config)
	fmt.Println(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.NewServer(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
