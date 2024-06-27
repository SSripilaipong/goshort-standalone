package access

import (
	"net/http"

	"goshort-standalone/common/fn"
	"goshort-standalone/common/httprsp"
)

func NewHandler(logic Logic) func(request *http.Request) httprsp.Writer {
	return fn.Compose3(httpResponse, logic.Access, parseRequest)
}
