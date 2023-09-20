package flower

import (
	"errors"
	"testing"
)

func prepareEnv() *Env {
	env := NewEnv(nil)
	env.Set("add", NewFuncValue(func(env *Env, args []Value) (Value, error) {
		var err error

		// Evaluate arguments before get
		args, err = env.EvalAll(args)
		if err != nil {
			return nil, err
		}

		// Get int values and add them together
		ok, i1 := GetIntValue(args[0])
		if !ok {
			return nil, errors.New("??")
		}
		ok, i2 := GetIntValue(args[1])
		if !ok {
			return nil, errors.New("??")
		}

		return NewIntValue(i1 + i2), nil
	}))
	env.Set("sum", NewFuncValue(func(env *Env, args []Value) (Value, error) {
		var err error

		// Make children env
		env2 := env.Fork()

		// Function x that returns 100
		env2.Set("x", NewFuncValue(func(env *Env, args []Value) (Value, error) {
			return NewIntValue(100), nil
		}))

		// Eval all arguments
		args, err = env2.EvalAll(args)
		if err != nil {
			return nil, err
		}

		// Let's add all ints together
		var summ int64 = 0
		for _, arg := range args {
			ok, i := GetIntValue(arg)
			if !ok {
				continue
			}
			summ += i
		}

		// And return it
		return NewIntValue(summ), nil

	}))
	return env
}

func prepareCall(atom string, args ...Value) *ExpValue {
	newArgs := []Value{NewAtomValue(atom)}
	newArgs = append(newArgs, args...)
	return NewExpValue(newArgs)
}

func TestEnvParentGet(t *testing.T) {
	env := NewEnv(nil)
	env.Set("A", NewIntValue(123))

	env2 := NewEnv(env)
	ok, i := GetIntValue(env2.Get("A"))
	if !ok {
		t.Fatal("Didn't found value for 'A'")
	}
	if i != 123 {
		t.Fatalf("Not correct %d", i)
	}
}

func TestEnvCallFunc(t *testing.T) {
	env := prepareEnv()
	exp := prepareCall("add", NewIntValue(32), NewIntValue(10))

	result, err := env.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	ok, i := GetIntValue(result)
	if !ok {
		t.Fatal("Not a correct type")
	}

	if i != 42 {
		t.Fatalf("Value is not 42: %d", i)
	}
}

func TestEnvCallChildren(t *testing.T) {
	env := prepareEnv()
	exp := prepareCall("sum", NewIntValue(200), prepareCall("x"), prepareCall("x"), NewIntValue(50))

	res, err := env.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	ok, i := GetIntValue(res)
	if !ok {
		t.Fatal("Not correct type")
	}

	if i != 450 {
		t.Fatalf("Wrong value. Not 450: %d", i)
	}

	// Then call that "x" will not be called outside of "sum"
	_, err = env.Eval(prepareCall("x"))
	if err == nil {
		t.Fatal("Should be an error as 'x' should not exist")
	}
}
