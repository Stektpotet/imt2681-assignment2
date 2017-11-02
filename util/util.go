package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

func GetEnv(env string) (value string) {
	value = os.Getenv(env)
	if value == "" {
		log.Fatalf("$%s must be set. See config.", env)
	}
	return
}

func DateString(y int, m time.Month, d int) string {
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// UntilTomorrow ...
func UntilTomorrow() time.Duration {
	// @doc https://stackoverflow.com/a/36988882
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	tomorrow = time.Date(
		tomorrow.Year(),
		tomorrow.Month(),
		tomorrow.Day(),
		0, 0, 0, 0,
		tomorrow.Location()) // Round to 00:00:00
	diff := tomorrow.Sub(now)

	// @debug
	log.Println("Tommorrow :", tomorrow)
	log.Println("Now       :", now)
	log.Println("Diff      :", diff)
	return diff
}
