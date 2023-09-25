package rt

func headTail[T any](arr []T, defval T) (T, []T) {
	if len(arr) < 1 {
		return defval, []T{}
	}
	if len(arr) < 2 {
		return arr[0], []T{}
	}
	return arr[0], arr[1:]
}
