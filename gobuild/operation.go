package gobuild

type GoOperation interface {

	// Renders the operation
	Render() (string, error)
}

// ================
// Operations
// ================

type GoNewVarOperation struct {
	// Variable name and type to be declared
	Info     *GoArgument
	DefValue GoValue
}
