// Package handlers contains implemented discordgo handlers
package handlers

import (
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

// MessageHandler handles incoming messages from discord.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	slog.Debug("Received message", slog.String("user", m.Author.Username), slog.String("guild", m.GuildID), slog.Int64("timestamp", time.Now().UnixMilli()))

	urls := xurls.Strict().FindAllString(m.Message.Content, -1)

	for _, url := range urls {
		slog.Debug("Found url in message", slog.String("url", url))
	}
}
