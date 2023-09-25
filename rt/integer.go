package rt

import "fmt"

type Integer struct {
	info   string
	number int64
}

func (self *Integer) Info() string {
	return self.info
}

func (self *Integer) String() string {
	return fmt.Sprint(self.number)
}

func (self *Integer) Int() int64 {
	return self.number
}

func (self *Integer) Float() float64 {
	return float64(self.number)
}

func NewInteger(info string, number int64) *Integer {
	return &Integer{
		info:   info,
		number: number,
	}
}
