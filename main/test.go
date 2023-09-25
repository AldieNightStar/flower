package main

import (
	"fmt"

	"github.com/AldieNightStar/flower"
	"github.com/AldieNightStar/flower/rt"
)

func main() {
	env := flower.NewEnv(nil)
	env.SetFunc("print", rt.ExpFunc(func(env *rt.Env, args []rt.Expression) (rt.Expression, error) {
		r, err := env.EvalAll(args)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
		return rt.NONE, nil
	}))
	result, err := flower.EvalString(env, "file", "(print (print (print 1 2) 3) 2 3)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
