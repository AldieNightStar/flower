package parser

import "testing"

func TestLexWord(t *testing.T) {
	w := lexWord("this ")
	if w != "this" {
		t.Fatalf("Wrong: %s", w)
	}

	w = lexWord("@@@")
	if w != "" {
		t.Fatalf("Wrong! Should be empty: %s", w)
	}
}

func TestLexDigits(t *testing.T) {
	d := lexDigits("12.34::")
	if d != "12.34" {
		t.Fatalf("Should be 12.34: %s", d)
	}

	d = lexDigits("23 1")
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
}

func TestLexSymbols(t *testing.T) {
	s := lexSymbols("@-->)")
	if s != "@-->" {
		t.Fatalf("Not correct: %s", s)
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
