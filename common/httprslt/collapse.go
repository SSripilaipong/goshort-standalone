package httprslt

import "goshort-standalone/common/httprsp"

func Collapse[A, B any](f func(A) B, g func(httprsp.Writer) B) func(Of[A]) B {
	return func(x Of[A]) B {
		if x.IsOk() {
			return f(x.Value())
		}
		return g(x.Error())
	}
}
