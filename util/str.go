package util

import "strings"

func Tabulate(s string) string {
	return "\t" + strings.ReplaceAll(s, "\n", "\n\t")
}

func SliceOf(s string, start, end int) string {
	if len(s) < 1 {
		return ""
	}
	LEN := len(s)
	if start < 0 {
		start = 0
	}
	if start > LEN {
		start = LEN
	}
	if end < start {
		end = start
	}
	if end > LEN {
		end = LEN
	}
	return s[start:end]
}

func SliceFrom(s string, start int) string {
	return SliceOf(s, start, 0xFFFFFFFF)
}

func StringAt(s string, index int) byte {
	if index < 0 {
		return 0
	} else if index >= len(s) {
		return 0
	}
	return s[index]
}
