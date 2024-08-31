package main

import (
	"github.com/devgugga/CyaniteMusic/config"
	"github.com/devgugga/CyaniteMusic/internal/bot"
	"log"
)

func main() {
	config.LoadConfig()

	discordBot, err := bot.NewBot(config.Cfg.Discord.Token)
	if err != nil {
		log.Fatalf("Failed to create Discord bot: %v", err)
	}

	discordBot.Start()
}
