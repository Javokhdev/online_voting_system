package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	VOTING_SERVICE_HOST string
	VOTING_SERVICE_PORT string

	PUBLIC_SERVICE_HOST string
	PUBLIC_SERVICE_PORT string

	HTTP_PORT string

}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.VOTING_SERVICE_HOST = cast.ToString(coalesce("VOTING_SERVICE_HOST", "localhost"))
	config.VOTING_SERVICE_PORT = cast.ToString(coalesce("VOTING_SERVICE_PORT", ":50051"))

	config.PUBLIC_SERVICE_HOST = cast.ToString(coalesce("PUBLIC_SERVICE_HOST", "localhost"))
	config.PUBLIC_SERVICE_PORT = cast.ToString(coalesce("PUBLIC_SERVICE_PORT", ":50050"))

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