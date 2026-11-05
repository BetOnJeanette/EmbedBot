// Package embeds creates the EmbedFixer interface, which should be implemented
// by all of the embeds that we replace for.
package embeds

import "net/url"

// EmbedFixer interface for all fixers
type EmbedFixer interface {
	Fix(link url.URL) string
	GetHost() string
}
