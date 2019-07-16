# Discord

A go library to quickly send events to discord channels

# Usage

1. Create a Webhook on the server, noting down the webhook URL.
2. In your Go application: `discord.WebhookURL = "<your webhook URL>"`
3. Now you can use `discord.Say("Whatever **you** _want_")`