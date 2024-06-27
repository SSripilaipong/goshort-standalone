package generate

import (
	"net/http"

	"goshort-standalone/common/fn"
	"goshort-standalone/common/httprsp"
)

func NewHandler(logic Logic) func(request *http.Request) httprsp.Writer {
	return fn.Compose(asHttpResponse(generate(logic)), parseRequest)
}
