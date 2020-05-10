package utils

import (
	"github.com/JeremyLoy/config"
	"log"
)

type EnvConfig struct {
	DEBUG       string
	HTTP_SERVER string
	HTTP_PORT   int
	DB_ADAPTER  string
	DB_HOST     string
	DB_PORT     int
	DB_DIALECT  string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

var ec EnvConfig

func GetEnvConfig() *EnvConfig {
	err := config.FromEnv().To(&ec)
	if err != nil {
		log.Fatal(err)
	}
	return &ec
}
