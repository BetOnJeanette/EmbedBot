// Package embeds creates the EmbedFixer interface, which should be implemented
// by all of the embeds that we replace for.
package embeds

type EmbedFixer interface {
	Fix(link string) string
	GetHost() string
}
