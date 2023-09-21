package parser

import "strings"

func ReadPath(s string) []string {
	return strings.Split(s, ".")
}
