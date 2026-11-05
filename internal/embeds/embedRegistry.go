package embeds

// EmbedRegistry holds a map of active embed fixers.
type EmbedRegistry struct {
	fixers map[string]EmbedFixer
}

// New registry with these commands
func New(fixers ...EmbedFixer) (registry *EmbedRegistry) {
	registry = &EmbedRegistry{fixers: make(map[string]EmbedFixer, len(fixers))}

	for _, fixer := range fixers {
		registry.fixers[fixer.GetHost()] = fixer
	}

	return
}

// Get fixer by hostname
func (registry *EmbedRegistry) Get(hostname string) EmbedFixer {
	if fixer, ok := registry.fixers[hostname]; ok {
		return fixer
	}

	return nil
}
