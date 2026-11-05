// Package fixers provides all the fixers for the various embeds we cover
package fixers

import (
	"net/url"
	"strings"
)

// Twitter embed fixer
type Twitter struct{}

// Fix Twitters embeds
func (Twitter) Fix(s url.URL) string {
	s.Host = "www.fixupx.com"
	q := s.Query()
	q.Del("a")
	q.Del("s")
	s.RawQuery = q.Encode()
	return s.String()
}

// GetHost for Twitter
func (Twitter) GetHost() string {
	return "x.com"
}

// ShouldFix links based on Twitter
func (Twitter) ShouldFix(link url.URL) bool {
	minSlashes := len(strings.Split("/realsupermari00/status/2115796218467295486", "/"))
	domainSplit := strings.Split(link.Path, "/")
	if (len(domainSplit) < minSlashes) { return false }
	return strings.ToLower(domainSplit[2]) == "status"
}
