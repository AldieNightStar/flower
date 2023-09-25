package rt

import "fmt"

type Word struct {
	info string
	name string
}

func (self *Word) Info() string {
	return self.info
}

func (self *Word) String() string {
	return fmt.Sprintf("Word: '%s'", self.name)
}

func (self *Word) Word() string {
	return self.name
}

func NewWord(info string, name string) *Word {
	return &Word{
		info: info,
		name: name,
	}
}
