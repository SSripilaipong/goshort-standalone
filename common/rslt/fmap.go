package rslt

func Fmap[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsOk() {
			return Value(f(x.Value()))
		}
		return Error[B](x.Error())
	}
}
