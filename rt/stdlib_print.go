package rt

import "fmt"

func stdlibPrint(env *Env, args []Expression) (Expression, error) {
	all, err := env.EvalAll(args)
	if err != nil {
		return NONE, err
	}
	for _, element := range all {
		fmt.Println(element)
	}
	return NONE, nil
}
