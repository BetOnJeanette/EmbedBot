// Package embedfixer creates the EmbedFixer interface, which should be implemented
// by all of the embeds that we replace for.
package embedfixer

type EmbedFixer interface {
	fix(link string) string
};
