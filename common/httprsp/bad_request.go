package httprsp

import "net/http"

func BadRequest(dataWriter DataWriter) WriterFunc {
	return func(writer http.ResponseWriter) error {
		if err := dataWriter.SetHeader(writer); err != nil {
			return err
		}
		writer.WriteHeader(http.StatusBadRequest)
		if err := dataWriter.WriteBody(writer); err != nil {
			return err
		}
		return nil
	}
}
