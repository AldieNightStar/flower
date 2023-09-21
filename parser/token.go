package parser

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
	TOK_EXP
	TOK_PATH
	TOK_COMMENT
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(t TokenType, value string) *Token {
	return &Token{t, value}
}
