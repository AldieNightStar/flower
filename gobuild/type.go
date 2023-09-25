package gobuild

type GoType struct {
	Name     string
	Package  string
	Generics []string
}

type GoArgument struct {
	Name string
	Type *GoType
}

type GoGeneric struct {
	Name       string
	Constraint string
}
