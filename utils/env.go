package utils

import (
	"fmt"
	"os"
)

var logger = NewLogger()

// MustGetEnv ...
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Error(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return value
}

// GetEnvOrDefault ...
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
