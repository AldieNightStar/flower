package parser

type Expression struct {
	Value     *Token
	Arguments []*Expression
}

func NewExpressionValue(value *Token) *Expression {
	return &Expression{
		Value:     value,
		Arguments: nil,
	}
}

func NewExpression(children []*Expression) *Expression {
	return &Expression{
		Value:     nil,
		Arguments: children,
	}
}
