package fixers

import "net/url"
type Reddit struct{}

func (Reddit) Fix(s url.URL) string {
	s.Host = "www.vxreddit.com"
	return s.String()
}

func (Reddit) GetHost() string {
	return "reddit.com"
}
