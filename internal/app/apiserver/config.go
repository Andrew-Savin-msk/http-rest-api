package apiserver

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{}
}

func ConfigLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading local.env file whith error:", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path is not set")
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatal("config file does not exists %s", configPath)
	}

	var cfg Config
	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
