package httprslt

import "goshort-standalone/common/fn"

func Fmap[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsOk() {
			return Value(f(x.Value()))
		}
		return Error[B](x.Error())
	}
}

func Join[T any](x Of[Of[T]]) Of[T] {
	if x.IsErr() {
		return Error[T](x.Error())
	}
	return x.Value()
}

func JointFmap[A, B any](f func(A) Of[B]) func(Of[A]) Of[B] {
	return fn.Compose(Join[B], Fmap(f))
}
