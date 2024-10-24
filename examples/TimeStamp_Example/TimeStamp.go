package main

import (
	"log"
	"time"

	"github.com/gtuk/discordwebhook"
)

func main() {
	var username = "BotUser"
	var content = "This is a test message"
	var url = "https://discord.com/api/webhooks/..."
	var image_url = "https://i.imgur.com/..."

	currentTime := time.Now()

	image := discordwebhook.Image{
		Url: &image_url,
	}

	embed := discordwebhook.Embed{
		Image:     &image,
		Timestamp: &currentTime,
	}

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
		Embeds:   &[]discordwebhook.Embed{embed},
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
