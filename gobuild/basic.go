package gobuild

type GoFunc struct {
	Name      string
	Arguments []*GoArgument
	Returns   []*GoArgument
	Body      []GoOperation
}

type GoStruct struct {
	Name   string
	Fields []*GoArgument
}

type GoFile struct {
	Package   string
	Functions map[string]*GoFunc
	Structs   map[string]*GoStruct
}
