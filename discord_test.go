package discord_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ecnepsnai/discord"
)

var webhookURL = ""

func TestMain(m *testing.M) {
	webhookURL = os.Getenv("DISCORD_WEBHOOK_URL")

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

func TestFileUpload(t *testing.T) {
	discord.WebhookURL = webhookURL
	f, err := os.OpenFile(path.Join(".", "discord.go"), os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Error opening file: %s", err.Error())
	}
	defer f.Close()
	content := discord.PostOptions{
		Content: "Hello, world!",
	}
	fileOptions := discord.FileOptions{
		FileName: "discord.go",
		Reader:   f,
	}
	if err := discord.UploadFile(content, fileOptions); err != nil {
		t.Errorf("Error posting message with file attachment: %s", err.Error())
	}
}
