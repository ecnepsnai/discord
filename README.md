# Discord

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/discord?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/discord)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/ecnepsnai/discord)
[![Releases](https://img.shields.io/github/release/ecnepsnai/discord/all.svg?style=flat-square)](https://github.com/ecnepsnai/discord/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/discord.svg?style=flat-square)](https://github.com/ecnepsnai/discord/blob/master/LICENSE)

A small go library to post messages to a Discord channel using Webhooks.

# Setup

You must first configure a Webhook on a Discord server before you can use this package. Instructions can be found on [Discord's support website](https://support.discord.com/hc/en-us/articles/228383668).

# Usage

You must first configure the Webhook URL that will be used

```golang
discord.WebhookURL = "https://discord.com/api/webhooks/.../..."
```

Then you can either send a simple text message or a more complex message

## Simple Text Message

```golang
discord.Say("Hello, world!")
```

## Complex Message

```golang
discord.Post(discord.PostOptions{
	Content: "Hello, world!",
	Embeds: []discord.Embed{
		{
			Author: discord.Author{
				Name: "ecnepsnai",
				URL:  "https://github.com/ecnepsnai",
			},
			Title:       "Amazing!",
			Description: "This is a cool embed",
		},
	},
})
```

## File Attachment

Restrictions with Discords Webhook API only supports 1 file upload at 8MiB or less.

```golang
var f *io.Reader // Pretend we've opened a file
content := discord.PostOptions{
	Content: "Hello, world!",
}
fileOptions := discord.FileOptions{
	FileName: "my_hot_mixtape.mp3",
	Reader:   f,
}
discord.UploadFile(content, fileOptions)
```

# Documentation

For more information see the [package's documentation](https://pkg.go.dev/github.com/ecnepsnai/discord).

**This package is not endorsed by or affiliated with Discord, inc.**
