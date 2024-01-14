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
	meta, err := toml.DecodeFile(configPath, config)
	fmt.Println(meta, err)
	fmt.Println(*config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
