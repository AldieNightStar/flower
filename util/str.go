package util

func SliceOf(s string, start, end int) string {
	LEN := len(s)
	if start < 0 {
		start = 0
	}
	if start > LEN {
		start = LEN
	}
	if end < start {
		end = start
	}
	if end > LEN {
		end = LEN
	}
	return s[start:end]
}
