package env

import (
	"os"
	"strconv"
)

func GetInt(key string) int {
	intVal, _ := strconv.Atoi(os.Getenv(key))
	return intVal
}

func GetInt64(key string) int64 {
	int64Val, _ := strconv.ParseInt(os.Getenv(key), 0, 0)
	return int64Val
}
