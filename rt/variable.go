package rt

type Variable interface {
	GetValue() Expression
}

type VariableMut interface {
	Variable
	SetValue(Expression) Expression
}

// ===================
// In memory variable
// ===================

type InMemVariable struct {
	Value Expression
}

func (v *InMemVariable) GetValue() Expression {
	return v.Value
}

func (v *InMemVariable) SetValue(val Expression) Expression {
	v.Value = val
	return val
}

func NewInMemVar(value Expression) *InMemVariable {
	return &InMemVariable{Value: value}
}
