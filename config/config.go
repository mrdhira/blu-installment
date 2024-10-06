package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
}

func readEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func ReadEnvInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func New() *Config {
	return &Config{
		DBHost: readEnv("DB_HOST", "localhost"),
		DBPort: ReadEnvInt("DB_PORT", 3306),
		DBUser: readEnv("DB_USER", "root"),
		DBPass: readEnv("DB_PASS", "root"),
		DBName: readEnv("DB_NAME", "blu_db"),
	}
}
