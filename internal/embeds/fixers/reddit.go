// Package fixers provides all the fixers for the various embeds we cover
package fixers

import "net/url"
// Reddit embed fixer
type Reddit struct{}

// Fix Reddits embeds
func (Reddit) Fix(s url.URL) string {
	s.Host = "www.vxreddit.com"
	return s.String()
}

// GetHost for Reddit
func (Reddit) GetHost() string {
	return "reddit.com"
}
