package rt

import "fmt"

type Atom struct {
	info  string
	value string
}

func (self *Atom) Info() string {
	return self.info
}

func (self *Atom) String() string {
	return fmt.Sprintf("\"%s\"", self.value)
}

func (self *Atom) AtomValue() string {
	return self.value
}

func (self *Atom) StringValue() string {
	return self.value
}

func NewAtom(info, value string) *Atom {
	return &Atom{
		info:  info,
		value: value,
	}
}
