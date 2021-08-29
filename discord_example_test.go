package discord_test

import (
	"io"

	"github.com/ecnepsnai/discord"
)

func ExampleSay() {
	discord.WebhookURL = "https://discord.com/api/webhooks/.../..."
	discord.Say("Hello, world!")
}

func ExamplePost() {
	discord.WebhookURL = "https://discord.com/api/webhooks/.../..."
	discord.Post(discord.PostOptions{
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
}

func ExampleUploadFile() {
	discord.WebhookURL = "https://discord.com/api/webhooks/.../..."
	var f io.Reader // Pretend we've opened a file
	content := discord.PostOptions{
		Content: "Hello, world!",
	}
	fileOptions := discord.FileOptions{
		FileName: "my_hot_mixtape.mp3",
		Reader:   f,
	}
	discord.UploadFile(content, fileOptions)
}
