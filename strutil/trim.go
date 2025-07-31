package strutil

import (
	"regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`^[^a-z0-9]+|[^a-z0-9]+$`)

func TrimSpec(s string) string {
	return nonAlphanumericRegex.ReplaceAllString(s, "")
}
