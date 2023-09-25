package rt

import "fmt"

type Float struct {
	info   string
	number float64
}

func (self *Float) Info() string {
	return self.info
}

func (self *Float) String() string {
	return fmt.Sprint(self.number)
}

func (self *Float) Float() float64 {
	return self.number
}

func (self *Float) Int() int64 {
	return int64(self.number)
}

func NewFloat(info string, value float64) *Float {
	return &Float{
		info:   info,
		number: value,
	}
}
