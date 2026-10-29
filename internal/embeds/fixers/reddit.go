// Package fixers provides all the fixers for the various embeds we cover
package fixers

import (
	"net/url"
	"strings"
)

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

// ShouldFix links based on Reddit
func (Reddit) ShouldFix(link url.URL) bool {
	domainSplit := strings.Split(link.Path, "/")
	if (len(domainSplit) < 6) { return false }
	return strings.ToLower(domainSplit[3]) == "comments"
}
