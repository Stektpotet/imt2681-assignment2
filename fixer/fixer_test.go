package fixer

import (
	"reflect"
	"testing"
)

func TestGetLatest(t *testing.T) {

	type args struct {
		constraints string
	}
	tests := []struct {
		name        string
		args        args
		wantPayload Currency
		wantErr     bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPayload, err := GetLatest(tt.args.constraints)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("GetLatest() = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}

func TestGetCurrencies(t *testing.T) {
	type args struct {
		constraints string
	}
	tests := []struct {
		name        string
		args        args
		wantPayload Currency
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPayload := GetCurrencies(tt.args.constraints); !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("GetCurrencies() = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}
