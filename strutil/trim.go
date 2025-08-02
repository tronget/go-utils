package strutil

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`^[^a-z0-9]+|[^a-z0-9]+$`)

func TrimSpec(s string) string {
	return nonAlphanumericRegex.ReplaceAllString(s, "")
}

func TrimSubstring(s, substr string) string {
	s = TrimSubLeft(s, substr)
	s = TrimSubRight(s, substr)
	return s
}

func TrimSubLeft(s, substr string) string {
	if substr == "" {
		return s
	}
	for strings.HasPrefix(s, substr) {
		s = s[len(substr):]
	}
	return s
}

func TrimSubRight(s, substr string) string {
	if substr == "" {
		return s
	}
	for strings.HasSuffix(s, substr) {
		s = s[:len(s)-len(substr)]
	}
	return s
}

func TrimToLength(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) > n {
		return s[:n]
	}
	return s
}
