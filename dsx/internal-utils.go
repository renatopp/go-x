package dsx

func resolveIndex(i, len int) int {
	if i < 0 {
		i += len
	}
	if i > len {
		i = len
	}
	return i
}
