package getkeyorcreate

import (
	"fmt"

	"goshort-standalone/common/rslt"
)

func New() func(url string) rslt.Of[string] {
	return func(url string) rslt.Of[string] {
		return rslt.Value(fmt.Sprintf("hash(%s)", url))
	}
}
