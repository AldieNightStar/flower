package parser

import "testing"

func TestLex(t *testing.T) {
	toks := Lex("", "word  :atom -12.44 'Hello'")

	if len(toks) != 7 {
		t.Fatalf("Wrong len: %d", len(toks))
	}

	if toks[0].Type != TOK_WORD {
		t.Fatalf("Wrong type")
	}
	if toks[1].Type != TOK_SPACE {
		t.Fatalf("Wrong type")
	}
	if toks[2].Type != TOK_ATOM {
		t.Fatalf("Wrong type")
	}
	if toks[3].Type != TOK_SPACE {
		t.Fatalf("Wrong type")
	}
	if toks[4].Type != TOK_NUMBER {
		t.Fatalf("Wrong type")
	}
	if toks[5].Type != TOK_SPACE {
		t.Fatalf("Wrong type")
	}
	if toks[6].Type != TOK_STRING {
		t.Fatalf("Wrong type")
	}
}

func TestLexLines(t *testing.T) {
	toks := Lex("file", "w1\nw2\n\nw4")

	if len(toks) < 5 {
		t.Fatalf("Len should be 5: %d", len(toks))
	}

	if toks[0].Info != "file:1" {
		t.Fatalf("Wrong: %s", toks[0].Info)
	}
	if toks[2].Info != "file:2" {
		t.Fatalf("Wrong: %s", toks[2].Info)
	}
	if toks[4].Info != "file:4" {
		t.Fatalf("Wrong: %s", toks[4].Info)
	}

	if toks[0].Value != "w1" {
		t.Fatalf("Wrong: %s", toks[0].Value)
	}
	if toks[2].Value != "w2" {
		t.Fatalf("Wrong: %s", toks[2].Value)
	}
	if toks[4].Value != "w4" {
		t.Fatalf("Wrong: %s", toks[4].Value)
	}
}

func TestLexOne(t *testing.T) {
	tok := lexOne("word  :atom -12.44 'Hello'", "", 1)

	if tok == nil {
		t.Fatal("nil")
	}

	if tok.Type != TOK_WORD {
		t.Fatalf("Not word: %d", tok.Type)
	}

	if tok.Value != "word" {
		t.Fatalf("Not word: %s", tok.Value)
	}
}

func TestLexWord(t *testing.T) {
	w := lexWord("this ")
	if w != "this" {
		t.Fatalf("Wrong: %s", w)
	}

	w = lexWord("@@@")
	if w != "" {
		t.Fatalf("Wrong! Should be empty: %s", w)
	}

	w = lexWord("a_b_c+")
	if w != "a_b_c" {
		t.Fatalf("Wrong: %s", w)
	}

	w = lexWord("tell-me!")
	if w != "tell-me" {
		t.Fatalf("Wrong: %s", w)
	}

	w = lexWord("-minus")
	if w != "" {
		t.Fatal("Should be empty")
	}

	w = lexWord("nil?")
	if w != "nil?" {
		t.Fatalf("Wrong: %s", w)
	}

	w = lexWord("?xyz")
	if w != "" {
		t.Fatal("Should be empty")
	}

}

func TestLexDigits(t *testing.T) {
	d := lexDigits("12.34::")
	if d != "12.34" {
		t.Fatalf("Should be 12.34: %s", d)
	}

	d = lexDigits("23-1")
	if d != "23" {
		t.Fatalf("Should be 23: %s", d)
	}

	d = lexDigits("1.2.3")
	if d != "1.2" {
		t.Fatalf("Should be 1.2: %s", d)
	}

	d = lexDigits(" 1")
	if d != "" {
		t.Fatalf("Should be empty: %s", d)
	}

	d = lexDigits("-44.38")
	if d != "-44.38" {
		t.Fatalf("Wrong: %s", d)
	}
}

func TestLexSymbols(t *testing.T) {
	s := lexSymbols("@-->)")
	if s != "@-->" {
		t.Fatalf("Not correct: %s", s)
	}

	s = lexSymbols("!**abc")
	if s != "!**" {
		t.Fatalf("Should be !**: %s", s)
	}

	s = lexSymbols(" 1")
	if s != "" {
		t.Fatalf("Should be empty: %s", s)
	}
}

func TestLexBracket(t *testing.T) {
	b := lexBracket("()")
	if b != "(" {
		t.Fatalf("Not correct: %s", b)
	}

	b = lexBracket(" [ )(")
	if b != "" {
		t.Fatalf("Should be empty: %s", b)
	}
}

func TestLexSpaces(t *testing.T) {
	s := lexSpaces("  \t\n1 ")
	if s != "  \t\n" {
		t.Fatalf("Not correct: %s.", s)
	}

	if lexSpaces("1") != "" {
		t.Fatal("Should be empty")
	}

	if takeEols(s) != 1 {
		t.Fatalf("Should be 1 eol: %d", takeEols(s))
	}
}

func TestLexString(t *testing.T) {
	s := lexString("'test' 1")
	if s != "'test'" {
		t.Fatalf("Wrong: %s", s)
	}

	s = lexString("'take\\'me'@@")
	if s != "'take\\'me'" {
		t.Fatalf("Wrong: %s", s)
	}

	if lexString("1 'a'") != "" {
		t.Fatal("Should be empty")
	}

	s = lexString("\"1\\\"\"")
	if s != "\"1\\\"\"" {
		t.Fatalf("Wrong: %s", s)
	}
}

func TestLexColonString(t *testing.T) {
	s := lexColonString(":hello 1")
	if s != ":hello" {
		t.Fatalf("Wrong: %s", s)
	}

	s = lexColonString(": 1")
	if s != "" {
		t.Fatalf("Should be empty: %s", s)
	}

	s = lexColonString(":a 1")
	if s != ":a" {
		t.Fatalf("Should be :a : %s", s)
	}

	s = lexColonString(" :1 2 3")
	if s != "" {
		t.Fatal("Should be empty")
	}
}

func TestLexComment(t *testing.T) {
	c := lexCommentString("; until end\n1")
	if c != "; until end\n" {
		t.Fatalf("Not correct: %s", c)
	}

	c = lexCommentString(";\n")
	if c != ";\n" {
		t.Fatalf("Not correct: %s", c)
	}

	c = lexCommentString("123")
	if c != "" {
		t.Fatal("Should be empty")
	}
}

func TestPathString(t *testing.T) {
	p := lexPathString("a.b.c")
	if p != "a.b.c" {
		t.Fatalf("Not correct: %s", p)
	}

	p = lexPathString("word and all")
	if p != "" {
		t.Fatalf("Should be empty: %s", p)
	}

	p = lexPathString("a.c==")
	if p != "a.c" {
		t.Fatalf("Not correct: %s", p)
	}

	p = lexPathString("a. b. c")
	if p != "" {
		t.Fatalf("Should be empty: %s", p)
	}

	p = lexPathString("no")
	if p != "" {
		t.Fatalf("Should be empty: %s", p)
	}

	p = lexPathString("a..b")
	if p != "" {
		t.Fatalf("Should be empty: %s", p)
	}
}
