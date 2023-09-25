package rt

import (
	"strings"

	"github.com/AldieNightStar/flower/util"
)

type Container struct {
	info     string
	children []Expression
}

func (self *Container) Info() string {
	return self.info
}

func (self *Container) String() string {
	sb := []string{}
	first := true
	for _, exp := range self.children {
		if first {
			sb = append(sb, exp.String())
			first = false
		} else {
			sb = append(sb, util.Tabulate(exp.String()))
		}
	}
	return "(" + strings.Join(sb, "\n") + ")"
}

func (self *Container) Children() []Expression {
	return self.children
}

func NewContainer(info string, children []Expression) *Container {
	return &Container{
		info:     info,
		children: children,
	}
}
