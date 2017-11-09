package webhook

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func PostHere(url, contentType string, r io.Reader) (resp *http.Response, err error) {
	resp = &http.Response{}
	resp.StatusCode = http.StatusOK
	resp.Header = http.Header{}
	resp.Header.Add("contentType", "application/json")
	raw, err := ioutil.ReadFile("../samples/base.json")
	if err != nil {
		log.Fatal(err) //could not read from file
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
	return
}

func TestSubsciptionOut_Invoke(t *testing.T) {

	r, err := PostHere("local", "application/json", nil)
	if err != nil {
		t.Error(err)
	}

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
		name        string
		fields      fields
		currentRate float32
		wantResp    *http.Response
		wantErr     bool
	}{
		{
			name: "OK",
			fields: fields{
				URL:    "local",
				Base:   "NOK",
				Target: "EUR",
				Min:    0.2,
				Max:    1,
			},
			currentRate: 0.4,
			wantResp:    r,
			wantErr:     false,
		},
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
			gotResp, err := hook.Invoke(tt.currentRate, PostHere)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubsciptionOut.Invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp == tt.wantResp {
				t.Errorf("SubsciptionOut.Invoke() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
