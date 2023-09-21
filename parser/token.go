package parser

type TokenType uint8

const (
	TOK_NONE TokenType = iota
	TOK_NUMBER
	TOK_STRING
	TOK_ATOM
	TOK_EXP
)

type Token struct {
	Type  TokenType
	Value string
}
