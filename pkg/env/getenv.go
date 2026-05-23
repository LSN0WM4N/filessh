package env

import (
	"os"
	"strconv"
)

// Return the `key` env, if this env does no exists return
// the value of fallback
func Getenv(key string, fallback string) string {
	value, exist := os.LookupEnv(key)

	if !exist {
		return fallback
	}

	return value
}

// `Getenv` implementation with Int parsing value
func GetIntEnv(key string, fallback int) int {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	result, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return result
}
