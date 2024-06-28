package httptst

import (
	"net/http"
)

func HandlerServer(pattern string, handler http.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(pattern, handler)
	return mux
}
