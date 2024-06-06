package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string
	TCP_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	PUB_DB_HOST     string
	PUB_DB_PORT     int
	PUB_DB_USER     string
	PUB_DB_PASSWORD string
	PUB_DB_NAME     string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	config.TCP_PORT = cast.ToString(coalesce("TCP_PORT", ":50051"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "voting"))

	config.PUB_DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.PUB_DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.PUB_DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.PUB_DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.PUB_DB_NAME = cast.ToString(coalesce("DB_NAME", "public"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}