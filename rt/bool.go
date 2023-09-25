package rt

import "fmt"

type Bool struct {
	info string
	flag bool
}

func (self *Bool) Info() string {
	return self.info
}

func (self *Bool) String() string {
	return fmt.Sprintf("\"%t\"", self.flag)
}

func (self *Bool) BoolValue() bool {
	return self.flag
}

func NewBool(info string, flag bool) *Bool {
	return &Bool{
		info: info,
		flag: flag,
	}
}
