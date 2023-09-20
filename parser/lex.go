package parser

import "strings"

var SPACES = " \t"
var SYMBOLS = "~!@#%^&*_+-=[]{}<>';\":\\/`,.?"
var SPACES_SYMBOLS = SPACES + SYMBOLS
var DIGIT = "01234567890"

func lexWord(src string) string {
	count := 0
	for _, c := range src {
		if strings.Contains(SPACES_SYMBOLS, string(c)) {
			break
		}
		count += 1
	}
	return src[0:count]
}

func lexDigits(src string) string {
	count := 0
	dotPresent := false
	for _, c := range src {
		// Allow for single dot
		if c == '.' {
			if dotPresent {
				// If dot is already present
				break
			} else {
				dotPresent = true
				count += 1
				continue
			}
		}
		if !strings.Contains(DIGIT, string(c)) {
			break
		}
		count += 1
	}
	return src[0:count]
}

func lexSymbols(src string) string {
	count := 0
	for _, c := range src {
		if !strings.Contains(SYMBOLS, string(c)) {
			break
		}
		count += 1
	}
	return src[0:count]
}

func lexBracket(src string) string {
	c := src[0:1]
	if c == "(" || c == ")" {
		return c
	}
	return ""
}
