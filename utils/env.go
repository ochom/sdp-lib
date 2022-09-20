package utils

import (
	"fmt"
	"os"
)

var logger = NewLogger()

// MustGetEnv ...
func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logger.Error(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return val
}

// GetEnvOrDefault ...
func GetEnvOrDefault(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}
