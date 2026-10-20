// Binary embed-bot.
package main

import (
	"os"

	"github.com/BetOnJeanette/EmbedBot/internal/bot"
)

func main() {
	token := os.Getenv("DISCORD_AUTH_TOKEN")

	bot, err := bot.New(token)

	if err != nil {
		os.Exit(1)
	}

	err = bot.Connect()

	if err != nil {
		os.Exit(2)
	}

	bot.Block()
	bot.Close()
}
