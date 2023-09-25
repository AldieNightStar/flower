package main

import (
	"fmt"

	"github.com/AldieNightStar/flower"
	"github.com/AldieNightStar/flower/rt"
)

func main() {
	env := flower.NewEnv(nil)
	env.SetFunc("print", rt.ExpFunc(func(args []rt.Expression) (rt.Expression, error) {
		fmt.Println(args)
		return rt.NONE, nil
	}))
	result, err := flower.EvalString(env, "file", "(print 1 2 3)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
