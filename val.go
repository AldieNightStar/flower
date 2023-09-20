package flower

import "fmt"

type Value interface {
	String() string
}

// =====================
// Exact value
// =====================

type ExactValue struct {
	Value any
}

func (this *ExactValue) String() string {
	if IsNoneValue(this) {
		return "NONE"
	}
	return fmt.Sprint(this.Value)
}

// =====================
// Number values
// =====================

func NewFloatValue(num float64) *ExactValue {
	return &ExactValue{
		Value: num,
	}
}

func NewIntValue(num int64) *ExactValue {
	return &ExactValue{
		Value: num,
	}
}

func GetIntValue(v Value) (ok bool, r int64) {
	exact, isExact := v.(*ExactValue)
	if !isExact {
		return false, 0
	}
	num, ok := exact.Value.(int64)
	if !ok {
		return false, 0
	}
	return true, num
}

func GetFloatValue(v Value) (ok bool, r float64) {
	exact, isExact := v.(*ExactValue)
	if !isExact {
		return false, 0
	}
	num, ok := exact.Value.(float64)
	if !ok {
		return false, 0
	}
	return true, num
}

// =====================
// String values
// =====================

func NewStringValue(s string) *ExactValue {
	return &ExactValue{
		Value: s,
	}
}

func GetStringValue(v Value) (ok bool, r string) {
	exact, isExact := v.(*ExactValue)
	if !isExact {
		return false, ""
	}
	s, ok := exact.Value.(string)
	if !ok {
		return false, ""
	}
	return true, s
}

// =====================
// Array values
// =====================

func NewArrayValue() *ExactValue {
	return &ExactValue{
		Value: make([]Value, 0, 8),
	}
}

func GetArrayValue(v Value) (ok bool, r []Value) {
	exact, isExact := v.(*ExactValue)
	if !isExact {
		return false, nil
	}
	arr, ok := exact.Value.([]Value)
	if !ok {
		return false, nil
	}
	return true, arr
}

// =====================
// Dict values
// =====================

func NewDictValue() Value {
	return &ExactValue{
		Value: make(map[string]Value, 32),
	}
}

func GetDictValue(v Value) (ok bool, r map[string]Value) {
	exact, isExact := v.(*ExactValue)
	if !isExact {
		return false, nil
	}
	dict, ok := exact.Value.(map[string]Value)
	if !ok {
		return false, nil
	}
	return true, dict
}

// =====================
// None Value
// =====================

var VALUE_NONE = &ExactValue{nil}

func IsNoneValue(v Value) bool {
	return v == VALUE_NONE
}

// =====================
// Func Value
// =====================

type FuncType func(env *Env, args []Value) (Value, error)

type FuncValue struct {
	Func FuncType
}

func (this *FuncValue) String() string {
	return "FUNC"
}

func GetFuncValue(v Value) (ok bool, f *FuncValue) {
	fun, isFunc := v.(*FuncValue)
	if !isFunc {
		return false, nil
	}
	return true, fun
}

func NewFuncValue(f FuncType) *FuncValue {
	return &FuncValue{f}
}

// =====================
// Atom Value
// =====================

type AtomValue struct {
	Atom string
}

func (this *AtomValue) String() string {
	return ":" + this.Atom
}

func NewAtomValue(atom string) *AtomValue {
	return &AtomValue{atom}
}

func GetAtomValue(v Value) (ok bool, a string) {
	atom, isAtom := v.(*AtomValue)
	if !isAtom {
		return false, ""
	}
	return true, atom.Atom
}

// =====================
// Expression Value
// =====================

type ExpValue struct {
	Args []Value
}

func (this *ExpValue) String() string {
	if len(this.Args) < 1 {
		return "()"
	} else {
		return "EXP(" + this.Args[0].String() + " ...)"
	}
}

func (this *ExpValue) Head() Value {
	if len(this.Args) < 1 {
		return VALUE_NONE
	} else {
		return this.Args[0]
	}
}

func (this *ExpValue) Tail() []Value {
	if len(this.Args) < 1 {
		return []Value{}
	} else {
		return this.Args[1:]
	}
}

func GetExpValue(v Value) (ok bool, e *ExpValue) {
	exp, isExp := v.(*ExpValue)
	if !isExp {
		return false, nil
	}
	return true, exp
}

func NewExpValue(args []Value) *ExpValue {
	return &ExpValue{args}
}
