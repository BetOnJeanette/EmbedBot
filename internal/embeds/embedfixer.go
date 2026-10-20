package embedfixer

type EmbedFixer interface {
	fix(link string) string
};
