package rt

type ExpressionAttributer interface {
	// Get attribute from the object
	// If no attribute then NONE should be returned
	GetAttribute(name string) Expression
}

type ExpressionMutableAttributer interface {
	// Set attribute to the object
	SetAttribute(name string, val Expression) Expression
}
