package discord_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ecnepsnai/discord"
)

var webhookURL = ""

func TestMain(m *testing.M) {
	for _, env := range os.Environ() {
		components := strings.Split(env, "=")
		key := components[0]
		value := components[1]
		if key == "DISCORD_WEBHOOK_URL" {
			webhookURL = value
			break
		}
	}

	if webhookURL == "" {
		fmt.Fprintf(os.Stderr, "DISCORD_WEBHOOK_URL environment variable is required\n")
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestSay(t *testing.T) {
	discord.WebhookURL = webhookURL
	err := discord.Say("Hello, world!")
	if err != nil {
		t.Errorf("Error posting plain-text message: %s", err.Error())
	}
}

func TestPost(t *testing.T) {
	discord.WebhookURL = webhookURL
	err := discord.Post(discord.PostOptions{
		Content: "Hello, world!",
		Embeds: []discord.Embed{
			{
				Color: 16777215,
				Author: &discord.Author{
					Name: "ecnepsnai",
					URL:  "https://github.com/ecnepsnai",
				},
				Title:       "Amazing!",
				Description: "This is a cool embed",
			},
		},
	})
	if err != nil {
		t.Errorf("Error posting complex message: %s", err.Error())
	}
}
