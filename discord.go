// Package discord is a go library to quickly send events to discord channels
// To get started: Create a Webhook on the server, noting down the webhook URL.
// Then, in your application set the webhook URL variable and then you can use `Say`
// for a simple text message, or `Post` for a more complex message.
package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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
	Author      *Author `json:"author,omitempty"`
	Title       string  `json:"title,omitempty"`
	URL         string  `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       uint32  `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Thumbnail   *Image  `json:"thumbnail,omitempty"`
	Image       *Image  `json:"image,omitempty"`
	Footer      *Footer `json:"footer,omitempty"`
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

// Image describes an image for an embed. If you need to upload an image
// you must use the `discord.UploadFile()` method, however that does not support
// rich embeds.
type Image struct {
	URL string `json:"url,omitempty"`
}

// Footer describes the footer for an embed
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

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(content); err != nil {
		return err
	}

	resp, err := http.Post(WebhookURL, "application/JSON", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}
	return nil
}

// FileOptions describes the options for uploading a file
type FileOptions struct {
	// The file name must include an extension and not include any directories
	FileName string
	Reader   io.Reader
}

// UploadFile will post a message to the channel for which the webhook is configured
// and attach the specified file to your message. Rich embeds are not supported and
// will be ignored if any are specified.
func UploadFile(content PostOptions, file FileOptions) error {
	if WebhookURL == "" {
		return nil
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fileWriter, err := w.CreateFormFile("file", file.FileName)
	if err != nil {
		return err
	}
	if _, err := io.Copy(fileWriter, file.Reader); err != nil {
		return err
	}
	payloadWriter, err := w.CreateFormField("payload_json")
	if err != nil {
		return err
	}
	if err := json.NewEncoder(payloadWriter).Encode(content); err != nil {
		return err
	}
	w.Close()

	req, err := http.NewRequest("POST", WebhookURL, &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}
	return nil
}
