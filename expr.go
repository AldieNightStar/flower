package flower

type Exp interface {
	Head() Exp
	Args() []Exp
	Eval() Exp
}

// ==============
// Static
// ==============
