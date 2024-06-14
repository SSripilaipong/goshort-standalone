package fn

func GoFn[A, B any](f func(A) B) func(A) {
	return func(x A) {
		go f(x)
	}
}
