package embedfixer
import "strings"

type reddit struct {}

func (reddit) fix(s string) string {
	return strings.Replace(s, "https://www.reddit.com/r", "https://www.vxreddit.com/r", 0);
}
