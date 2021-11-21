package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	JWTSecretKey string
}

func NewConfig() *Config {
	config := Config{}

	config.JWTSecretKey = cast.ToString(getOrReturnDefault("JWT_SECRET", "this_is_secret_key"))
	return &config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
