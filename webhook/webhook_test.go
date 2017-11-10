package webhook

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math"
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

	body := &RequestBody{
		Base:    "NOK",
		Target:  "EUR",
		Current: 1.5,
		Min:     1,
		Max:     2,
	}

	tests := []struct {
		name        string
		currentRate float32
		wantInvoke  bool
	}{
		{
			name:        "Too low -> does not invoke",
			currentRate: 0.5,
			wantInvoke:  true,
		},
		{
			name:        "In Range -> does invoke",
			currentRate: 1.5,
			wantInvoke:  true,
		},
		{
			name:        "Too high -> does not invoke",
			currentRate: 10,
			wantInvoke:  false,
		},
		{
			name:        "negative value -> does not invoke",
			currentRate: -1,
			wantInvoke:  false,
		},
		{
			name:        "value -> does not invoke",
			currentRate: float32(math.NaN()),
			wantInvoke:  false,
		},
	}
	for _, tt := range tests {
		body.Current = tt.currentRate
		t.Run(tt.name, func(t *testing.T) {
			hook := &SubsciptionOut{
				Base:   body.Base,
				Target: body.Target,
				Min:    body.Min,
				Max:    body.Max,
			}
			gotBody := hook.Invoke(tt.currentRate, PostHere)
			if gotBody == body {
				t.Errorf("SubsciptionOut.Invoke() = %v, want %v", gotBody, body)
			}
		})
	}
}
