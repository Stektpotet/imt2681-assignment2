package webhook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SubsciptionIn struct {
	HookID string  `json:"hookID,omitempty"`
	URL    string  `json:"webhookURL"`
	Base   string  `json:"baseCurrency"`
	Target string  `json:"targetCurrency"`
	Min    float32 `json:"minTriggerValue"`
	Max    float32 `json:"maxTriggerValue"`
}

type SubsciptionOut struct {
	URL    string  `json:"webhookURL"`
	Base   string  `json:"baseCurrency"`
	Target string  `json:"targetCurrency"`
	Min    float32 `json:"minTriggerValue"`
	Max    float32 `json:"maxTriggerValue"`
}

type WebhookRequestBody struct {
	Base    string  `json:"baseCurrency"`
	Target  string  `json:"targetCurrency"`
	Current float32 `json:"currentRate"`
	Min     float32 `json:"minTriggerValue"`
	Max     float32 `json:"maxTriggerValue"`
}

func (hook *SubsciptionOut) Invoke(currentRate float32, client http.Client) (resp *http.Response, err error) {
	body := WebhookRequestBody{
		hook.Base,
		hook.Target,
		currentRate,
		hook.Min,
		hook.Max,
	}
	raw, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = client.Post(hook.URL, "application/json", bytes.NewBuffer(raw))
	return
}
