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

func NewFloatValue(num float64) Value {
	return &ExactValue{
		Value: num,
	}
}

func NewIntValue(num int64) Value {
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

func NewStringValue(s string) Value {
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

func NewArrayValue() Value {
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

type FuncValue struct {
	Func func(args []Value) Value
}

func (this *FuncValue) String() string {
	return "FUNC"
}

func NewFuncValue(f func(args []Value) Value) Value {
	return &FuncValue{f}
}
