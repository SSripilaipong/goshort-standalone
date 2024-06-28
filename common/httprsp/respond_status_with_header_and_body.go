package httprsp

import "net/http"

func respondStatusWithHeaderAndBody(status int) func(dataWriter DataWriter) WriterFunc {
	return func(dataWriter DataWriter) WriterFunc {
		return func(writer http.ResponseWriter) error {
			if err := dataWriter.SetHeader(writer); err != nil {
				return err
			}
			writer.WriteHeader(status)
			if err := dataWriter.WriteBody(writer); err != nil {
				return err
			}
			return nil
		}
	}
}
