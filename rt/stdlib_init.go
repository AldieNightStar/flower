package rt

func AddStdLib(env *Env) {
	// func (env *Env, args []Expression) (Expression, error)

	addStdLibConsole(env)
}

func addStdLibConsole(env *Env) {
	// func (env *Env, args []Expression) (Expression, error)

	env.SetFunc("print", stdlibPrint)
}

func addStdLibLogic(env *Env) {
	// func (env *Env, args []Expression) (Expression, error)
	env.SetFunc("=", stdlibEq)
}
