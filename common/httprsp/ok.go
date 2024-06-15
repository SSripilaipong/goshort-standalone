package httprsp

import "net/http"

func Ok(dataWriter DataWriter) WriterFunc {
	return func(writer http.ResponseWriter) error {
		if err := dataWriter.SetHeader(writer); err != nil {
			return err
		}
		writer.WriteHeader(http.StatusOK)
		if err := dataWriter.WriteBody(writer); err != nil {
			return err
		}
		return nil
	}
}
