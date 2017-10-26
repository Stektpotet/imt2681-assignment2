package main

import "fmt"

//WebhookPayload - The payload of webhooks in the system
type WebhookPayload struct {
	URL             string  `json:"webhookURL"`
	Base            string  `json:"baseCurrency"`
	Target          string  `json:"targetCurrency"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

func SanetizeHook(hookURL string) string {
	return fmt.Sprintf("Sanetization not implemented. Hook:\n%+v", hookURL)
}

func main() {

}
