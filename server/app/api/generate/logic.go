package generate

import "goshort-standalone/common/rslt"

type Logic interface {
	Generate(url string) rslt.Of[string]
}

type LogicFunc func(url string) rslt.Of[string]

func (f LogicFunc) Generate(url string) rslt.Of[string] {
	return f(url)
}
