/*
Command disco provides a command-line interface to post Discord messages to a channel
using a familiar interface like sendmail.

To use disco you must first define the webhook in any of three places:

 1. Populate the `DISCORD_WEBHOOK_URL` environment variable with your webhook URL
 2. Create a file at `~/.discord` with the webhook URL
 3. Create a file at `/etc/discord` with the webhook URL

disco will look for the webhook URL in that order, allowing you to define a fallback webhook
for the entire system, while letting users override that and use thier own.

disco reads from stdin for the message and only supports text messages, it does not
support sending attachments.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/ecnepsnai/discord"
)

func main() {
	u, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "No webhook URL defined\n")
		os.Exit(1)
	}

	discord.WebhookURL = *u

	if err := discord.Say(getStdin()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func getStdin() string {
	text := ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return text
}

func loadConfig() (*string, error) {
	// 1. Look for the environment variable DISCORD_WEBHOOK_URL
	if url := os.Getenv("DISCORD_WEBHOOK_URL"); url != "" {
		return &url, nil
	}

	// 2. Try to read ~/.discord
	if webhookURL, _ := readConfigFile(path.Join(os.Getenv("HOME"), ".discord")); webhookURL != nil {
		return webhookURL, nil
	}

	// 3. Try to read /etc/discord
	return readConfigFile(path.Join("/", "etc", "discord"))
}

func readConfigFile(filePath string) (*string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	webhookURL := string(data)

	// Trim newline
	if webhookURL[len(webhookURL)-1] == '\n' {
		webhookURL = webhookURL[0 : len(webhookURL)-1]
	}

	return &webhookURL, nil
}
