package generate

import (
	"goshort-standalone/common/fn"
	"goshort-standalone/common/httprslt"
	"goshort-standalone/common/httprsp"
	"goshort-standalone/common/rslt"
)

type Response struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func asHttpResponse(f func(JsonBody) rslt.Of[urlWithKey]) func(httprslt.Of[JsonBody]) httprsp.Writer {
	return httprslt.Collapse(fn.Compose(genResultToHttpResponse, f), fn.Id[httprsp.Writer])
}

var genResultToHttpResponse = rslt.Collapse(func(result urlWithKey) httprsp.Writer {
	return httprsp.Ok(httprsp.Json(Response{Url: result.Url, Key: result.Key}))
}, func(err error) httprsp.Writer {
	return httprsp.InternalServerError(httprsp.Json(httprsp.NewMessage(err.Error())))
})
