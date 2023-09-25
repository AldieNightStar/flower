package flower

import (
	"testing"

	"github.com/AldieNightStar/flower/parser"
	"github.com/AldieNightStar/flower/rt"
)

func TestConverter(t *testing.T) {
	r, err := ConvertRuntime(
		parser.NewTokenNode(parser.NewToken(
			parser.TOK_ATOM,
			"hello",
			"file",
			1,
		)),
	)

	if err != nil {
		t.Fatal(err)
	}

	atom, atomOk := r.(*rt.Atom)
	if !atomOk {
		t.Fatal("Should be atom")
	}

	if atom.AtomValue() != "hello" {
		t.Fatalf("Should be hello: %s", atom.AtomValue())
	}
}
