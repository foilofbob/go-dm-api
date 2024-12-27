package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Database struct {
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
}

var currentConfig = Config{
	Server: Server{
		Port: "123",
		Host: "localhost",
	},
	Database: Database{
		Username: "db",
		Password: "pass",
	},
}

func InitConfig() {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Println(fmt.Sprintf("Error opening config: %s", err.Error()))
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&currentConfig)
	if err != nil {
		log.Println(fmt.Sprintf("Error decoding config: %s", err.Error()))
	}
}

func Cfg() Config {
	return currentConfig
}
