// Package handlers contains implemented discordgo handlers
package handlers

import (
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/BetOnJeanette/EmbedBot/internal/embeds"
	"github.com/BetOnJeanette/EmbedBot/internal/embeds/fixers"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/publicsuffix"
	"mvdan.cc/xurls/v2"
)

// MessageHandler parses and handles incoming messages from discord.
type MessageHandler struct {
	registry *embeds.EmbedRegistry
}

// Init message handler
func Init() *MessageHandler {
	registry :=
		embeds.New(
			fixers.Reddit{},
		)

	return &MessageHandler{registry: registry}
}

// Handler handles incoming messages from discord.
func (messageHandler *MessageHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	slog.Debug("Received message", slog.String("user", m.Author.Username), slog.String("guild", m.GuildID), slog.Int64("timestamp", time.Now().UnixMilli()))

	links := xurls.Strict().FindAllString(m.Message.Content, -1)

	var sb strings.Builder

	for _, link := range links {
		u, err := url.Parse(link)

		if err != nil {
			slog.Error("Failed to parse link", slog.String("link", link))
			continue
		}

		host, err := publicsuffix.EffectiveTLDPlusOne(u.Hostname())

		if err != nil {
			slog.Error("Error parsing url hostname")
			continue
		}

		slog.Debug("Found link in message", slog.String("host", host), slog.String("scheme", u.Hostname()), slog.String("path", u.EscapedPath()))

		fixer := messageHandler.registry.Get(host)

		if fixer == nil {
			slog.Debug("No fixer found for host")
			continue
		}

		sb.WriteString("-# ")
		sb.WriteString(fixer.Fix(link))
		sb.WriteRune('\n')

	}

	if sb.Len() == 0 {
		return
	}

	s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		ID:      m.ID,
		Channel: m.ChannelID,
		Flags:   m.Flags | discordgo.MessageFlagsSuppressEmbeds,
	})

	_, err := s.ChannelMessageSend(m.ChannelID, sb.String())

	if err != nil {
		slog.Error("Error sending message to discord", slog.String("err", err.Error()))
	}
}
