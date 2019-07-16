// Package discord is a go library to quickly send events to discord channels
// To get started: Create a Webhook on the server, noting down the webhook URL.
// Then, in your application set the webhook URL variable and then you can use `Say`
// with whatever you want to say. Formatting is supported.
package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookURL your discord webhook URL
var WebhookURL string

// Say send a message to the discord channel
// If WebhookURL is not set, this does nothing
func Say(message string) error {
	if WebhookURL == "" {
		return nil
	}

	data, err := json.Marshal(map[string]string{
		"content": message,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(WebhookURL, "application/JSON", bytes.NewReader(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}
	return nil
}
