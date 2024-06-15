package httprsp

import (
	"net/http"
)

func InternalServerError(dataWriter DataWriter) WriterFunc {
	return func(writer http.ResponseWriter) error {
		if err := dataWriter.SetHeader(writer); err != nil {
			return err
		}
		writer.WriteHeader(http.StatusInternalServerError)
		if err := dataWriter.WriteBody(writer); err != nil {
			return err
		}
		return nil
	}
}
