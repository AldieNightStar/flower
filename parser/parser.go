package parser

import (
	"errors"
)

func Parse(tokens []*Token) ([]*Node, error) {
	nodeStack := []*Node{NewChildrenNode()}
	for _, tok := range tokens {
		// Some token types are going to be skipped
		if tok.Type == TOK_SPACE || tok.Type == TOK_COMMENT {
			continue
		}

		if tok.Type == TOK_BRACKET {
			if tok.Value == "(" {
				nodeStack = append(nodeStack, NewChildrenNode())
			} else if tok.Value == ")" {
				LEN := len(nodeStack)
				last := nodeStack[LEN-1]
				nodeStack[LEN-2].addsubnode(last)
				nodeStack = nodeStack[0 : len(nodeStack)-1]
			}
		} else {
			nodeStack[len(nodeStack)-1].addsubnode(NewTokenNode(tok))
		}
	}
	LEN := len(nodeStack)
	if LEN > 1 {
		return nil, errors.New("Too many open nodes")
	} else if LEN < 1 {
		return nil, errors.New("Redundant closing found")
	}
	return nodeStack[0].Children, nil
}
