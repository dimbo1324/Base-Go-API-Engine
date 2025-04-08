package env

import (
	"os"
	"strconv"
)

func GetString(key, fallBack string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallBack
	}
	return val
}
func GetInt(key string, fallBack int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallBack
	}
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallBack
	}
	return valAsInt
}
