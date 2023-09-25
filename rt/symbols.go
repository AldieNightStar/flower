package rt

import "fmt"

type Symbols struct {
	info  string
	value string
}

func (self *Symbols) Info() string {
	return self.info
}

func (self *Symbols) String() string {
	return fmt.Sprintf("Symbols '%s'", self.value)
}

func (self *Symbols) Symbols() string {
	return self.value
}

func NewSymbols(info string, value string) *Symbols {
	return &Symbols{
		info:  info,
		value: value,
	}
}
