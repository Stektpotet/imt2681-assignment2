package webhook

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func GetFromHere(url string) (resp *http.Response, err error) {
	resp = &http.Response{}
	resp.StatusCode = http.StatusOK
	resp.Header = http.Header{}
	resp.Header.Add("contentType", "application/json")
	raw, err := ioutil.ReadFile("../samples/base.json")
	log.Fatal(err)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
	return
}
func GetEmptyFromHere(url string) (resp *http.Response, err error) {
	resp = &http.Response{}
	resp.StatusCode = http.StatusNoContent
	return
}

func TestSubsciptionOut_Invoke(t *testing.T) {
	type fields struct {
		URL    string
		Base   string
		Target string
		Min    float32
		Max    float32
	}
	type args struct {
		currentRate float32
		client      http.Client
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
	// {"OK", fields{}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := &SubsciptionOut{
				URL:    tt.fields.URL,
				Base:   tt.fields.Base,
				Target: tt.fields.Target,
				Min:    tt.fields.Min,
				Max:    tt.fields.Max,
			}
			gotResp, err := hook.Invoke(tt.args.currentRate, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubsciptionOut.Invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SubsciptionOut.Invoke() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
