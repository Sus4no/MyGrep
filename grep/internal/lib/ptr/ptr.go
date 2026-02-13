package ptr

func ZeroIfNil[T any](ptr *T) T {
	if ptr == nil {
		var val T
		return val
	}
	return *ptr
}
