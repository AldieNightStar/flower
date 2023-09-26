package rt

type ComparableEq interface {
	CompareEq(ComparableEq) Expression
}
