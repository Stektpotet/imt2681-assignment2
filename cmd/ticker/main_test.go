package main

import (
	"testing"

	"github.com/stektpotet/imt2681-assignment2/fixer"
)

func Test_initializeDBConnection(t *testing.T) {
	type args struct {
		mongoDBHosts []string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializeDBConnection(tt.args.mongoDBHosts)
		})
	}
}
func TestInvokeHooks(t *testing.T) {
	type args struct {
		current fixer.Currency
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InvokeHooks(tt.args.current)
		})
	}
}
