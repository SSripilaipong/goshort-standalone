package geturl

import (
	"fmt"

	"goshort-standalone/common/rslt"
)

func New() func(key string) rslt.Of[string] {
	return func(key string) rslt.Of[string] {
		return rslt.Value(fmt.Sprintf("https://www.google.com/search?q=%s", key))
	}
}
