package access

import (
	"goshort-standalone/common/httprsp"
	"goshort-standalone/common/rslt"
)

var httpResponse = rslt.Collapse(func(url string) httprsp.Writer {
	return httprsp.PermanentRedirect(url)
}, func(err error) httprsp.Writer {
	return httprsp.InternalServerError(httprsp.Json(httprsp.NewMessage(err.Error())))
})
