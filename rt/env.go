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
		return f.Call(args)
	}
	return nil, errors.New("Not a function")
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

func (this *Env) Eval(e Expression) (Expression, error) {
	if e == NONE {
		return NONE, nil
	}
	// Words take values if there are
	word, isWord := e.(ExpressionWord)
	if isWord {
		return this.Get(word.Word()), nil
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
