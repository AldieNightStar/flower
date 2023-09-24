package parser

import "testing"

func TestParse(t *testing.T) {
	lexed := Lex("file", "(word 'string 123' (inside 2) :sym) (999)")
	parsed, err := Parse(lexed)
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 2 {
		t.Fatalf("Must be 2: %d", len(parsed))
	}

	// Tesint first node
	if parsed[0].IsValue() {
		t.Fatal("Cannot be value node")
	}
	if parsed[0].Children[0].Token.Type != TOK_WORD {
		t.Fatal("Should be a word")
	}
	if parsed[0].Children[1].Token.Type != TOK_STRING {
		t.Fatal("Should be a string")
	}
	if !parsed[0].Children[2].IsContainer() {
		t.Fatal("Should be a container")
	}
	if parsed[0].Children[2].Children[0].Token.Value != "inside" {
		t.Fatal("Should be value 'inside'")
	}
	if parsed[0].Children[2].Children[1].Token.Value != "2" {
		t.Fatal("Should be value '2'")
	}
	if parsed[0].Children[3].Token.Type != TOK_ATOM {
		t.Fatal("Should be an atom token")
	}

	// Testing second node
	if !parsed[1].IsContainer() {
		t.Fatal("Should be a container")
	}
	if parsed[1].Children[0].Token.Value != "999" {
		t.Fatal("Should be a '999' value")
	}
}
