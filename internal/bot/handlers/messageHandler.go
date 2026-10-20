// Package handlers contains implemented discordgo handlers
package handlers

import (
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler handles incoming messages from discord.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	slog.Debug("Received message", slog.String("user", m.Author.Username), slog.String("guild", m.GuildID), slog.Int64("timestamp", time.Now().UnixMilli()))
}
