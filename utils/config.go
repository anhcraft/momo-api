package utils

import (
	"fmt"
	"os"
)

func GetEnvSetting(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func RequireEnvSetting(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s env setting is absent", key))
	}
	return v
}
