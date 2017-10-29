package util

import (
	"log"
	"os"
)

func GetEnv(env string) (value string) {
	value = os.Getenv(env)
	if value == "" {
		log.Fatalf("$%s must be set. See config.", env)
	}
	return
}
