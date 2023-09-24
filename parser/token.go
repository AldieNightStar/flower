package parser

import "fmt"

type TokenType uint8

const (
	TOK_NONE TokenType = iota
	TOK_SPACE
	TOK_NUMBER
	TOK_STRING
	TOK_SYMBOLS
	TOK_ATOM
	TOK_WORD
	TOK_BRACKET
	TOK_PATH
	TOK_COMMENT
	TOK_NODE
)

type Token struct {
	Info  string
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token %s [%d] '%s'", t.Info, t.Type, t.Value)
}

func NewToken(t TokenType, value string, fileName string, line int) *Token {
	return &Token{
		Info:  fmt.Sprintf("%s:%d", fileName, line),
		Type:  t,
		Value: value,
	}
}
