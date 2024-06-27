package httprslt

import (
	"goshort-standalone/common/either"
	"goshort-standalone/common/httprsp"
)

type Of[T any] either.Of[httprsp.Writer, T]

func (e Of[T]) IsOk() bool {
	return e.either().IsRight()
}

func (e Of[T]) IsErr() bool {
	return e.either().IsLeft()
}

func (e Of[T]) either() either.Of[httprsp.Writer, T] {
	return either.Of[httprsp.Writer, T](e)
}

func (e Of[T]) Value() T {
	return e.either().Right()
}

func (e Of[T]) Error() httprsp.Writer {
	return e.either().Left()
}

func Value[T any](x T) Of[T] {
	return Of[T](either.Right[httprsp.Writer, T](x))
}

func Error[T any](err httprsp.Writer) Of[T] {
	return Of[T](either.Left[httprsp.Writer, T](err))
}

func ValueOf[T any](x Of[T]) T {
	return x.Value()
}
