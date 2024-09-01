package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/devgugga/CyaniteMusic.git/internal/bot/handlers"
	"log"
)

type Bot struct {
	Session *discordgo.Session
}

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
	b.Session.AddHandler(b.onInteractionCreate)
}

func (b *Bot) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		handlers.CommandHandler(s, i)
	}
}

func (b *Bot) Start() {
	if err := b.Session.Open(); err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}

	log.Println("Bot is now running. Press CTRL+C to stop")
	select {}
}

func (b *Bot) onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Printf("Bot connected as %s#%s", s.State.User.Username, s.State.User.Discriminator)

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}
}
