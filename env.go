package utils

import (
	"strconv"
	"strings"
)

func GetEnv(env map[string]string, key string, defaultVal string) string {
	if val, ok := env[key]; ok {
		return val
	}

	return defaultVal
}

func GetEnvAsInt(env map[string]string, name string, defaultVal int) int {
	valueStr := GetEnv(env, name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func GetEnvAsUInt(env map[string]string, name string, defaultVal uint) uint {
	valueStr := GetEnv(env, name, "")
	if value, err := strconv.ParseUint(valueStr, 10, 64); err == nil {
		return uint(value)
	}

	return defaultVal
}

func GetEnvAsBool(env map[string]string, name string, defaultVal bool) bool {
	valStr := GetEnv(env, name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsSlice(env map[string]string, name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(env, name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
