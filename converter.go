package flower

import (
	"strconv"
	"strings"

	"github.com/AldieNightStar/flower/parser"
	"github.com/AldieNightStar/flower/rt"
)

func ConvertRuntime(node *parser.Node) (rt.Expression, error) {
	if node.IsContainer() {
		children := []rt.Expression{}

		for _, child := range node.Children {
			expr, err := ConvertRuntime(child)
			if err != nil {
				return nil, err
			}
			children = append(children, expr)
		}

		// TODO add info
		rt.NewContainer("", children)
	} else if node.IsValue() {
		t := node.Token.Type
		if t == parser.TOK_NONE {
			return rt.NONE, nil
		} else if t == parser.TOK_NUMBER {
			if strings.Contains(node.Token.Value, ".") {
				// Float value
				f, err := strconv.ParseFloat(node.Token.Value, 64)
				if err != nil {
					return nil, err
				}
				return rt.NewFloat(node.Token.Info, f), nil
			} else {
				// Int value
				i, err := strconv.ParseInt(node.Token.Value, 10, 64)
				if err != nil {
					return nil, err
				}
				return rt.NewInteger(node.Token.Info, i), nil
			}
		} else if t == parser.TOK_STRING {
			return rt.NewStr(node.Token.Info, node.Token.Value), nil
		} else if t == parser.TOK_SYMBOLS {
			return rt.NewSymbols(node.Token.Info, node.Token.Value), nil
		} else if t == parser.TOK_ATOM {
			return rt.NewAtom(node.Token.Info, node.Token.Value), nil
		} else if t == parser.TOK_WORD {
			return rt.NewWord(node.Token.Info, node.Token.Value), nil
		} else if t == parser.TOK_PATH {
			return rt.NewPath(node.Token.Info, parser.ReadPath(node.Token.Value)), nil
		}
	}
	return rt.NONE, nil
}

func ConvertRuntimeAll(nodes []*parser.Node) ([]rt.Expression, error) {
	arr := []rt.Expression{}
	for _, node := range nodes {
		e, err := ConvertRuntime(node)
		if err != nil {
			return nil, err
		}
		arr = append(arr, e)
	}
	return arr, nil
}
