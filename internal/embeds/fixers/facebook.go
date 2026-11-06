// Package fixers provides all the fixers for the various embeds we cover
package fixers

import (
	"net/url"
)

// Facebook embed fixer
type Facebook struct{}

// Fix Facebooks embeds
func (Facebook) Fix(s url.URL) string {
	s.Host = "www.facebed.com"
	return s.String()
}

// GetHost for Facebook
func (Facebook) GetHost() string {
	return "facebook.com"
}

// ShouldFix links based on Facebook
func (Facebook) ShouldFix(_ url.URL) bool {
	return true
}
