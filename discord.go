// Package discord is a go library to quickly send events to discord channels
// To get started: Create a Webhook on the server, noting down the webhook URL.
// Then, in your application set the webhook URL variable and then you can use `Say`
// for a simple text message, or `Post` for a more complex message.
package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookURL your discord webhook URL
var WebhookURL string

// Say sends the provided message to the channel for which the webhook is configured.
// If WebhookURL is not set, this does nothing.
func Say(message string) error {
	return Post(PostOptions{Content: message})
}

// PostOptions describes all possible options for a post
type PostOptions struct {
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

// Embed describes embedded content within a message
type Embed struct {
	Author      Author  `json:"author,omitempty"`
	Title       string  `json:"title,omitempty"`
	URL         string  `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       uint32  `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Thumbnail   Image   `json:"thumbnail,omitempty"`
	Image       Image   `json:"image,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
}

// Author describes the author for an embed
type Author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Field describes a field for an embed
type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// Image describes an image for an embed
type Image struct {
	URL string `json:"url,omitempty"`
}

// Footer describes the foorter for an embed
type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Post will post a message to the channel for which the webhook is configured.
// Unlike `discord.Say()`, Post gives you full control over the message.
func Post(content PostOptions) error {
	if WebhookURL == "" {
		return nil
	}

	data, err := json.Marshal(content)
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
