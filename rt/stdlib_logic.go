package rt

import "fmt"

func stdlibEq(env *Env, args []Expression) (Expression, error) {
	if len(args) < 2 {
		return NONE, fmt.Errorf("Need at least two arguments")
	}
	evaled, err := env.EvalAll(args)
	if err != nil {
		return NONE, err
	}
	c1, iscomp1 := evaled[0].(ComparableEq)
	if !iscomp1 {
		return NewBool("", false), nil
	}
	c2, iscomp2 := evaled[0].(ComparableEq)
	if !iscomp2 {
		return NewBool("", false), nil
	}
	return c1.CompareEq(c2), nil
}
