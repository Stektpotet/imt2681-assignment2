package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

//GetEnv - Obtain environment variable value
func GetEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Fatalf("$%s must be set. See config.", key)
	}
	return
}

//DateString - get a given date in the following format: yyyy-mm-dd
func DateString(y int, m time.Month, d int) string {
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

//Contains - looks for a given string within the string array
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Benchmark(msg string, start time.Time) {
	duration := (time.Now().Nanosecond() - start.Nanosecond()) / 1e6
	fmt.Printf("Duration for %s: %vms\n", msg, duration)
}
