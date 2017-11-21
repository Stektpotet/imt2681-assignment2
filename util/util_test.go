package util

import (
	"log"
	"testing"
	"time"
)

func TestGetEnv(t *testing.T) {
	logFatalf = func(format string, v ...interface{}) { log.Printf(format, v) }
	defer func() { logFatalf = log.Fatalf }()
	if !t.Run("Fail Please", func(t *testing.T) {
		GetEnv("NOTHING")
	}) {
		t.Error("Expected fail with nonexistant env variable")
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Containing item", args{[]string{"localhost", "someotherhost", "anotherhost.com"}, "localhost"}, true},
		{"Does not contain item", args{[]string{"localhost", "someotherhost", "anotherhost.com"}, "youtube.com"}, false},
		{"Sorta' contains item", args{[]string{"localhoster", "someotherhost", "anotherhost.com"}, "localhost"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

//As DateString does not handle 'invalid' dates, as it is always used in cohesion with time.Time.Date()
func TestDateString(t *testing.T) {
	type args struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Before October(i.e. 01-09)", args{2004, time.September, 24}, "2004-09-24"},
		{"After October(i.e. 10-12)", args{2024, time.November, 10}, "2024-11-10"},
		{"Before The tenth", args{2004, time.September, 4}, "2004-09-04"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateString(tt.args.y, tt.args.m, tt.args.d); got != tt.want {
				t.Errorf("DateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
