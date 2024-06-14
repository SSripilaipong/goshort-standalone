package slc

func Apply[T any](f func(T), xs []T) {
	for _, x := range xs {
		f(x)
	}
}
