package parser

import "testing"

func TestReadString(t *testing.T) {
	var r string

	r = ReadString("'this is the string'")
	if r != "this is the string" {
		t.Fatalf("Incorrect: %s", r)
	}

	r = ReadString("'error string")
	if r != "" {
		t.Fatalf("Should be empty: %s", r)
	}

	r = ReadString(":atom")
	if r != "atom" {
		t.Fatalf("Incorrect: %s", r)
	}

	r = ReadString("hello there")
	if r != "" {
		t.Fatalf("Should be empty: %s", r)
	}
}
