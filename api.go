package flower

import (
	"github.com/AldieNightStar/flower/parser"
	"github.com/AldieNightStar/flower/rt"
)

func NewEnv(parent *rt.Env) *rt.Env {
	return rt.NewEnv(parent)
}

func EvalString(env *rt.Env, filename string, src string) (rt.Expression, error) {
	nodes, err := parser.Parse(parser.Lex(filename, src))
	if err != nil {
		return nil, err
	}

	expressions, err := ConvertRuntimeAll(nodes)
	if err != nil {
		return nil, err
	}

	results, err := env.EvalAll(expressions)
	if err != nil {
		return nil, err
	}

	if len(results) < 0 {
		return rt.NONE, nil
	}

	return results[len(results)-1], nil
}
