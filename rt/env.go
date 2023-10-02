package rt

import "errors"

type Env struct {
	Parent *Env
	Scope  map[string]Expression
}

func NewEnv(parent *Env) *Env {
	return &Env{
		Parent: parent,
		Scope:  make(map[string]Expression, 32),
	}
}

func (this *Env) Fork() *Env {
	return NewEnv(this)
}

func (this *Env) Get(name string) Expression {
	val, ok := this.Scope[name]
	if !ok {
		if this.Parent != nil {
			return this.Parent.Get(name)
		} else {
			return NONE
		}
	}
	return val
}

func (this *Env) Set(name string, e Expression) Expression {
	this.Scope[name] = e
	return e
}

func (this *Env) CallFunc(e Expression, args []Expression) (Expression, error) {
	f, isFunc := e.(ExpressionFunction)
	if isFunc {
		return f.Call(this, args)
	}
	return nil, errors.New("Not a function")
}

func (this *Env) SetFunc(name string, fn ExpFunc) {
	this.Scope[name] = NewFuncExpression("native", fn)
}

func (this *Env) EvalAll(expressions []Expression) (res []Expression, err error) {
	for _, exp := range expressions {
		r, err := this.Eval(exp)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (this *Env) EvalPath(pathExp ExpressionPath) (res Expression, err error) {
	path := pathExp.Path()
	var e Expression
	for id, name := range path {
		if id < 1 {
			// If this is the first time then we take it from ENV
			e = this.Get(name)
		} else {
			// Next time we should take from inner values
			if e == NONE {
				return nil, errors.New("none can't have any attribute")
			}
			attr, ok := e.(ExpressionAttributer)
			if !ok {
				return nil, errors.New(e.String() + ": Has no attributes inside")
			}
			e = attr.GetAttribute(name)
		}
	}
	return e, nil
}

func (this *Env) Eval(e Expression) (Expression, error) {
	if e == NONE {
		return NONE, nil
	}

	// Words take values if there are
	word, isWord := e.(ExpressionWord)
	if isWord {
		return this.Get(word.Word()), nil
	}

	// Symbols take values if there are
	symb, isSym := e.(ExpressionSymbols)
	if isSym {
		return this.Get(symb.Symbols()), nil
	}

	path, isPath := e.(ExpressionPath)
	if isPath {
		return this.EvalPath(path)
	}
	// TODO add path words support
	// ...

	// If this is container then eval it
	container, isContainer := e.(ExpressionContainer)
	if isContainer {
		var err error
		head, tail := headTail(container.Children(), nil)
		head, err = this.Eval(head)
		if err != nil {
			return nil, err
		}
		// TODO Function need to know whos is callee
		return this.CallFunc(head, tail)
	}
	return e, nil
}
