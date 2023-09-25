package rt

type ExpFunc func(args []Expression) (Expression, error)

type FuncExpression struct {
	info string
	fn   ExpFunc
}

func NewFuncExpression(info string, fn ExpFunc) *FuncExpression {
	return &FuncExpression{
		info: info,
		fn:   fn,
	}
}

func (self *FuncExpression) String() string {
	return "(func ...)"
}

func (self *FuncExpression) Info() string {
	return self.info
}

func (self *FuncExpression) Call(args []Expression) (Expression, error) {
	return self.fn(args)
}
