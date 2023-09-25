package rt

import "fmt"

type Str struct {
	info  string
	value string
}

func (self *Str) Info() string {
	return self.info
}

func (self *Str) String() string {
	return fmt.Sprintf("\"%s\"", self.value)
}

func (self *Str) StringValue() string {
	return self.value
}

func NewStr(info, value string) *Str {
	return &Str{
		info:  info,
		value: value,
	}
}
