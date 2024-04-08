package apiserver

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func ConfigLoad() *Config {
	fmt.Println("Loading config")
	var configPathEnv = "CONFIG_PATH_DOCKER"
	if runtime.GOOS == "windows" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading local.env file whith error:", err)
		}
		configPathEnv = "CONFIG_PATH"
	}
	fmt.Println("Env path set")
	configPath := os.Getenv(configPathEnv)
	// configPath := os.Getenv("CONFIG_PATH_DOCKER")
	fmt.Println(configPath)
	if configPath == "" {
		log.Fatal("config path is not set")
	}
	fmt.Println("config loaded")
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file does not exists %s", configPath)
	}

	var cfg Config
	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
