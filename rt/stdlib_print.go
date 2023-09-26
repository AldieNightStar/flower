package rt

import "fmt"

func stdlibPrint(env *Env, args []Expression) (Expression, error) {
	all, err := env.EvalAll(args)
	if err != nil {
		return NONE, err
	}
	fmt.Println(all)
	if len(all) > 0 {
		return all[len(all)-1], nil
	}
	return NONE, nil
}
