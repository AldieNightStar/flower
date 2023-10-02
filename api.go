package flower

import (
	"github.com/AldieNightStar/flower/parser"
	"github.com/AldieNightStar/flower/rt"
	"github.com/AldieNightStar/flower/util"
)

func NewEnv(parent *rt.Env) *rt.Env {
	var env = rt.NewEnv(parent)
	rt.AddStdLib(env)
	return env
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

func EvalFile(env *rt.Env, filename string) (rt.Expression, error) {
	fileSrc, err := util.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return EvalString(env, filename, fileSrc)
}
