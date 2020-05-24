package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/ecnepsnai/discord"
)

func main() {
	u, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "No webhook URL defined\n")
		os.Exit(1)
	}

	discord.WebhookURL = *u

	discord.Say(getStdin())
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
	envMap := getEnvMap()
	if url, ok := envMap["DISCORD_WEBHOOK_URL"]; ok {
		return &url, nil
	}

	// 2. Try to read .discord
	if webhookURL, _ := readConfigFile(path.Join(envMap["HOME"], ".discord")); webhookURL != nil {
		return webhookURL, nil
	}

	// 3. Try to read /etc/discord
	return readConfigFile(path.Join("etc", "discord"))
}

func getEnvMap() map[string]string {
	m := map[string]string{}

	for _, s := range os.Environ() {
		k, v := kvSplit(s)
		m[k] = v
	}

	return m
}

func kvSplit(in string) (key string, value string) {
	components := strings.SplitN(in, "=", 2)
	key = components[0]
	value = components[1]
	return
}

func readConfigFile(filePath string) (*string, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	webhookURL := string(data)
	return &webhookURL, nil
}
