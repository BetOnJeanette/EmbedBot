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
	q := s.Query()
	q.Del("utm_source")
	q.Del("utm_name")
	q.Del("utm_term")
	q.Del("utm_content")
	q.Del("utm_medium")
	s.RawQuery = q.Encode()
	return s.String()
}

// GetHost for Reddit
func (Reddit) GetHost() string {
	return "reddit.com"
}

// ShouldFix links based on Reddit
func (Reddit) ShouldFix(link url.URL) bool {
	minSlashes := len(strings.Split("/r/THE_PACK/comments/1xg8hwi", "/"))
	domainSplit := strings.Split(link.Path, "/")
	if (len(domainSplit) < minSlashes) { return false }
	return strings.ToLower(domainSplit[3]) == "comments"
}
