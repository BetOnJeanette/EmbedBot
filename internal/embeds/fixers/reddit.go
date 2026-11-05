package fixers

import "regexp"

type Reddit struct{}

var redditRe = regexp.MustCompile(`(?i)^(https?://)(?:[a-z0-9-]+\.)*reddit\.com([:/?#]|$)`)

func (Reddit) Fix(s string) string {
	return redditRe.ReplaceAllString(s, "${1}vxreddit.com${2}")
}

func (Reddit) GetHost() string {
	return "reddit.com"
