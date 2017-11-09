package webhook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// SubsciptionIn - The webhook as stored within the system
type SubsciptionIn struct {
	HookID string  `json:"hookID,omitempty"`
	URL    string  `json:"webhookURL"`
	Base   string  `json:"baseCurrency"`
	Target string  `json:"targetCurrency"`
	Min    float32 `json:"minTriggerValue"`
	Max    float32 `json:"maxTriggerValue"`
}

// SubsciptionOut - The webhook as it looks "outside" the system
type SubsciptionOut struct {
	URL    string  `json:"webhookURL"`
	Base   string  `json:"baseCurrency"`
	Target string  `json:"targetCurrency"`
	Min    float32 `json:"minTriggerValue"`
	Max    float32 `json:"maxTriggerValue"`
}

// RequestBody - The body sent when invoking the webhook
type RequestBody struct {
	Base    string  `json:"baseCurrency"`
	Target  string  `json:"targetCurrency"`
	Current float32 `json:"currentRate"`
	Min     float32 `json:"minTriggerValue"`
	Max     float32 `json:"maxTriggerValue"`
}

type postFunc func(string, string, io.Reader) (*http.Response, error)

// Invoke - Invoke a post request to the webhook's URL with it's own body.
func (hook *SubsciptionOut) Invoke(currentRate float32, poster postFunc) (resp *http.Response, err error) {
	body := &RequestBody{
		Base:    hook.Base,
		Target:  hook.Target,
		Current: currentRate,
		Min:     hook.Min,
		Max:     hook.Max,
	}
	raw, _ := json.Marshal(body)
	resp, err = poster(hook.URL, "application/json", bytes.NewBuffer(raw))
	return
}
