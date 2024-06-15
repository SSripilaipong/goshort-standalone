package httprsp

import "net/http"

func PermanentRedirect(url string) WriterFunc {
	return func(writer http.ResponseWriter) error {
		writer.Header().Set("Location", url)
		writer.WriteHeader(http.StatusPermanentRedirect)
		return nil
	}
}
