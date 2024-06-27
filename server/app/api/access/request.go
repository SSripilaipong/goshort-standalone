package access

import "net/http"

func parseRequest(request *http.Request) string {
	return request.PathValue("key")
}
