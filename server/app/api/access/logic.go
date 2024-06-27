package access

import "goshort-standalone/common/rslt"

type Logic interface {
	Access(key string) rslt.Of[string]
}

type LogicFunc func(key string) rslt.Of[string]

func (f LogicFunc) Access(key string) rslt.Of[string] {
	return f(key)
}
