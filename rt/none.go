package rt

var NONE = None(0)

type None byte

func (self None) String() string {
	return "None"
}

func (self None) Info() string {
	return "none"
}
