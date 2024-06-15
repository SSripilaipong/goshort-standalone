package generate

import (
	"encoding/json"
	"net/http"

	"goshort-standalone/common/httprsp"
)

type JsonBody struct {
	Url string `json:"url"`
}

type Response struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func NewHandler(logic Logic) func(request *http.Request) httprsp.Writer {
	return func(request *http.Request) httprsp.Writer {
		decoder := json.NewDecoder(request.Body)
		var body JsonBody
		err := decoder.Decode(&body)
		if err != nil {
			return httprsp.BadRequest(httprsp.Json(httprsp.NewMessage("bad request")))
		}
		key, err := logic.Generate(body.Url)
		if err != nil {
			return httprsp.InternalServerError(httprsp.Json(httprsp.NewMessage(err.Error())))
		}
		return httprsp.Ok(httprsp.Json(Response{Url: body.Url, Key: key}))
	}
}
