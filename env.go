package flower

import "errors"

type Env struct {
	Parent *Env
	Scope  map[string]Value
}

func NewEnv(parent *Env) *Env {
	return &Env{
		Parent: parent,
		Scope:  make(map[string]Value, 32),
	}
}

func (this *Env) Fork() *Env {
	return NewEnv(this)
}

func (this *Env) Get(name string) Value {
	val, ok := this.Scope[name]
	if !ok {
		if this.Parent != nil {
			return this.Parent.Get(name)
		} else {
			return VALUE_NONE
		}
	}
	return val
}

func (this *Env) Set(name string, v Value) Value {
	this.Scope[name] = v
	return v
}

func (this *Env) CallFunc(v Value, args []Value) (Value, error) {
	f := GetFuncValue(v)
	if f != nil {
		return f.Func(this, args)
	}
	return nil, errors.New("Not a function")
}

func (this *Env) EvalAll(values []Value) (res []Value, err error) {
	for _, val := range values {
		r, err := this.Eval(val)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (this *Env) Eval(v Value) (Value, error) {
	isNone := IsNoneValue(v)
	if isNone {
		return VALUE_NONE, nil
	}
	isAtom, atom := GetAtomValue(v)
	if isAtom {
		return this.Get(atom), nil
	}
	isExp, exp := GetExpValue(v)
	if isExp {
		var err error
		head, err := this.Eval(exp.Head())
		if err != nil {
			return nil, err
		}
		args := exp.Tail()
		return this.CallFunc(head, args)
	}
	return v, nil
}
