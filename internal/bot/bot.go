package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
}

// NewBot creates a new Discord bot instance using the provided token.
// It initializes a new Discord session using the discordgo library and registers handlers.
//
// token: The bot token obtained from the Discord Developer Portal.
//
// Returns:
// - A pointer to a Bot struct representing the new bot instance.
// - An error if the bot initialization fails.
func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{Session: dg}
	bot.registerHandlers()

	return bot, nil
}

func (b *Bot) registerHandlers() {
	b.Session.AddHandler(b.onReady)
}

// Start starts the Discord bot session.
// It opens the connection to Discord using the session and logs a message indicating that the bot is running.
// The function blocks indefinitely until the bot is manually stopped.
//
// If the session fails to open due to an error, it logs the error message and fatally exits the program.
func (b *Bot) Start() {
	if err := b.Session.Open(); err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}

	log.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func (b *Bot) onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Printf("Bot connected as %s#%s", s.State.User.Username, s.State.User.Discriminator)
}
