package rt

import (
	"fmt"
	"strings"
)

type Path struct {
	info string
	path []string
}

func (self *Path) Info() string {
	return self.info
}

func (self *Path) String() string {
	return fmt.Sprintf("Path: '%s'", strings.Join(self.path, "."))
}

func (self *Path) Path() []string {
	return self.path
}

func NewPath(info string, path []string) *Path {
	return &Path{
		info: info,
		path: path,
	}
}
