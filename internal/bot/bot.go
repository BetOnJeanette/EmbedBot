// Package bot implements discordgo functionality for our project.
package bot

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/BetOnJeanette/EmbedBot/internal/bot/handlers"
	"github.com/bwmarrin/discordgo"
)

// Bot struct.
type Bot struct {
	dg *discordgo.Session
}

// New connects to discord using discordgo and return a newly created bot instance.
func New(token string) (bot *Bot, err error) {
	dg, err := discordgo.New(token)

	if err != nil {
		slog.Error("Couldn't initialize discordgo session.", "err", err)
		return
	}

	slog.Info("Initialized discordgo session.")

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentsGuildMessages

	dg.AddHandler(handlers.Init().Handler)

	bot = &Bot{dg}

	return
}

// Connect to discord.
func (bot *Bot) Connect() error {
	err := bot.dg.Open()

	if err != nil {
		slog.Error("Couldn't connect to discord.", "err", err)
		return err
	}

	slog.Info("Successfully connected to discord.")

	return nil
}

// Block until we receive an interrupt
func (bot *Bot) Block() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Close connection to discord.
func (bot *Bot) Close() {
	bot.dg.Close()

	slog.Info("Disconnected from discord.")
}
