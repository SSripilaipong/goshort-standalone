package fn

func Unvoid[B, A any](f func(A)) func(A) B {
	return func(x A) (y B) {
		f(x)
		return y
	}
}
