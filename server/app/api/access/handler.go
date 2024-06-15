package access

import (
	"net/http"

	"goshort-standalone/common/httprsp"
)

func NewHandler(logic Logic) func(request *http.Request) httprsp.Writer {
	return func(request *http.Request) httprsp.Writer {
		key := request.PathValue("key")

		url, err := logic.Access(key)

		switch {
		case err == nil:
			return httprsp.PermanentRedirect(url)
		default:
			return httprsp.InternalServerError(httprsp.Json(httprsp.NewMessage(err.Error())))
		}
	}
}
