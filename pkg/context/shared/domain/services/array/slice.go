package array

func Slice[T any](values []T, index int) (T, bool) {
	if index < 0 || index >= len(values) {
		return *new(T), false
	}

	return values[index], true
}
