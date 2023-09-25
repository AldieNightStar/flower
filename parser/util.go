package parser

import "strings"

func ReadPath(s string) []string {
	return strings.Split(s, ".")
}

func ReadString(s string) string {
	if len(s) < 2 {
		return ""
	}
	// First symbol quote should be equal to the last one
	if s[0] == '\'' || s[0] == '"' || s[0] == '`' {
		q := s[0]
		last := len(s) - 1
		if s[last] == q {
			return s[1:last]
		} else {
			return ""
		}
		// Or if it's an atom
	} else if s[0] == ':' {
		return s[1:]
	}
	return ""
}
