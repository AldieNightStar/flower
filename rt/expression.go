package rt

type Expression interface {
	Info() string
	String() string
}

type ExpressionContainer interface {
	Expression
	Children() []Expression
}

type ExpressionInteger interface {
	Expression
	Int() int64
}

type ExpressionFloat interface {
	Expression
	Float() float64
}

type ExpressionBoolean interface {
	Expression
	BoolValue() bool
}

type ExpressionString interface {
	Expression
	StringValue() string
}

type ExpressionAtom interface {
	Expression
	AtomValue() string
}

type ExpressionSymbols interface {
	Expression
	Symbols() string
}

type ExpressionPath interface {
	Expression
	Path() []string
}

type ExpressionWord interface {
	Expression
	Word() string
}

type ExpressionFunction interface {
	Expression
	Call([]Expression) (Expression, error)
}
