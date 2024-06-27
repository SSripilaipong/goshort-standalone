package generate

import (
	"encoding/json"
	"net/http"

	"goshort-standalone/common/httprslt"
	"goshort-standalone/common/httprsp"
)

type JsonBody struct {
	Url string `json:"url"`
}

func parseRequest(request *http.Request) httprslt.Of[JsonBody] {
	decoder := json.NewDecoder(request.Body)
	var body JsonBody
	if err := decoder.Decode(&body); err != nil {
		return httprslt.Error[JsonBody](httprsp.BadRequest(httprsp.Json(httprsp.NewMessage("bad request"))))
	}
	return httprslt.Value(body)
}
