package choose

func One[T any](condition bool, t, f T) T {
	if !condition {
		return f
	}

	return t
}
