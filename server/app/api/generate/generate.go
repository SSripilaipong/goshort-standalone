package generate

import (
	"goshort-standalone/common/rslt"
)

type urlWithKey struct {
	Url string
	Key string
}

func generate(logic Logic) func(JsonBody) rslt.Of[urlWithKey] {
	return func(body JsonBody) rslt.Of[urlWithKey] {
		key := logic.Generate(body.Url)
		return rslt.Fmap(combineUrlAndKey(body.Url))(key)
	}
}

var combineUrlAndKey = func(url string) func(key string) urlWithKey {
	return func(key string) urlWithKey {
		return urlWithKey{Url: url, Key: key}
	}
}
