package parser

import (
	"strings"

	"github.com/AldieNightStar/flower/util"
)

var SPACES = " \t"
var EOL = "\n"
var SPACES_EOL = SPACES + EOL
var SYMBOLS = "~!@#$%^&*_+-=[]{}<>';\":\\/`,.?"
var SPACES_SYMBOLS = SPACES + SYMBOLS
var SPACES_SYMBOLS_BRACKETS = SPACES + SYMBOLS + "()"
var DIGIT = "01234567890"
var QUOTES = "\"'`"

func Lex(src string) []*Token {
	var tokens []*Token

	pos := 0
	for {
		tok := lexOne(src[pos:])
		if tok == nil {
			break
		}
		tokens = append(tokens, tok)
		pos += len(tok.Value)
	}

	return tokens
}

func lexOne(src string) *Token {
	var res string

	res = lexSpaces(src)
	if len(res) > 0 {
		return NewToken(TOK_SPACE, res)
	}

	res = lexDigits(src)
	if len(res) > 0 {
		return NewToken(TOK_NUMBER, res)
	}

	res = lexPathString(src)
	if len(src) > 0 {
		return NewToken(TOK_PATH, res)
	}

	res = lexWord(src)
	if len(res) > 0 {
		return NewToken(TOK_WORD, res)
	}

	res = lexString(src)
	if len(res) > 0 {
		return NewToken(TOK_STRING, res)
	}

	res = lexBracket(src)
	if len(res) > 0 {
		return NewToken(TOK_BRACKET, res)
	}

	res = lexColonString(src)
	if len(res) > 0 {
		return NewToken(TOK_ATOM, res)
	}

	res = lexCommentString(src)
	if len(res) > 0 {
		return NewToken(TOK_COMMENT, res)
	}

	res = lexSymbols(src)
	if len(res) > 0 {
		return NewToken(TOK_SYMBOLS, res)
	}

	return nil
}

func lexWord(src string) string {
	count := 0
	for _, c := range src {
		if c == '$' || c == '_' {
			count += 1
			continue
		}
		if c == '-' {
			if count > 0 {
				count += 1
				continue
			} else {
				break
			}
		}
		if strings.Contains(SPACES_SYMBOLS_BRACKETS, string(c)) {
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
		if c == '-' {
			if count > 0 {
				break
			}
			count += 1
			continue
		}
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

func lexSpaces(src string) string {
	count := 0
	for _, c := range src {
		if !strings.Contains(SPACES_EOL, string(c)) {
			break
		}
		count += 1
	}
	return src[0:count]
}

func takeEols(s string) int {
	var eols = 0
	for _, c := range s {
		if c == '\n' {
			eols += 1
		}
	}
	return eols
}

func lexString(s string) string {
	if !strings.Contains(QUOTES, s[0:1]) {
		return ""
	}
	q := rune(s[0])
	count := 1
	escaped := false
	for _, c := range s[1:] {
		if escaped {
			escaped = false
			count += 1
			continue
		}
		if c == '\\' {
			escaped = true
			count += 1
			continue
		}
		if c == q {
			count += 1
			break
		}
		count += 1
	}
	return s[0:count]
}

func lexColonString(s string) string {
	if s[0] != ':' {
		return ""
	}
	count := 1 + len(lexWord(s[1:]))
	if count < 2 {
		return ""
	}
	return s[0:count]
}

func lexCommentString(s string) string {
	if s[0] != ';' {
		return ""
	}
	count := 1
	for _, c := range s[1:] {
		if c == '\n' {
			count += 1
			break
		}
		count += 1
	}
	return s[0:count]
}

func lexPathString(s string) string {
	count := 0
	delta := 0
	for {
		delta = len(lexWord(s[count:]))
		if delta > 0 {
			count += delta
		}
		if util.SliceOf(s, count, count+1) != "." {
			break
		} else {
			count += 1
		}
	}
	result := util.SliceOf(s, 0, count)
	if strings.Count(result, ".") > 0 {
		// If last dot is present then it's incorrect
		if strings.LastIndex(result, ".") == len(result)-1 {
			return ""
		}
		return result
	} else {
		return ""
	}
}
