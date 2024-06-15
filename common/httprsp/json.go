package httprsp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(payload interface{}) DataWriterImpl {
	return NewDataWriter(
		BodyWriterFunc(func(writer http.ResponseWriter) error {
			body, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("json marshal: %w", err)
			}
			_, err = writer.Write(body)
			return err
		}),
		HeaderSetterFunc(func(writer http.ResponseWriter) error {
			writer.Header().Set("Content-Type", "application/json")
			return nil
		}),
	)
}
