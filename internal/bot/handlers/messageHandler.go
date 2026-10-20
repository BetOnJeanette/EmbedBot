// Package handlers contains implemented discordgo handlers
package handlers

import (
	"log/slog"
	"net/url"
	"time"

	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

// MessageHandler handles incoming messages from discord.
func MessageHandler(_ *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	slog.Debug("Received message", slog.String("user", m.Author.Username), slog.String("guild", m.GuildID), slog.Int64("timestamp", time.Now().UnixMilli()))

	links := xurls.Strict().FindAllString(m.Message.Content, -1)

	for _, link := range links {
		u, err := url.Parse(link)

		if err != nil {
			slog.Error("Failed to parse link", slog.String("link", link))
			continue
		}

		slog.Debug("Found link in message", slog.String("host", u.Host), slog.String("scheme", u.Scheme), slog.String("path", u.EscapedPath()))
	}
}
