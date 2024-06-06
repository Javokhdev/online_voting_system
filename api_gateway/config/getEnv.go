package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	SERVICE_HOST string
	SERVICE_PORT string

	HTTP_PORT string

}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.SERVICE_HOST = cast.ToString(coalesce("SERVICE_HOST", "localhost"))
	config.SERVICE_PORT = cast.ToString(coalesce("SERVICE_PORT", ":50051"))
	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":7070"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}