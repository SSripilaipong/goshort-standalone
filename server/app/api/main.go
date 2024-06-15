package api

import (
	"net/http"

	"goshort-standalone/common/httprsp"
	"goshort-standalone/server/app/api/access"
	"goshort-standalone/server/app/api/generate"
)

func New(logic logic) *App {
	return newApp(8080, newApiHandler(logic))
}

func newApiHandler(logic logic) *http.ServeMux {
	mux := http.NewServeMux()
	route(mux, "POST /generate", generate.NewHandler(logic))
	route(mux, "GET /a/{key}", access.NewHandler(logic))
	return mux
}

func route(mux *http.ServeMux, pattern string, handler func(request *http.Request) httprsp.Writer) {
	mux.Handle(pattern, handlerFunc(handler))
}

func handlerFunc(f func(request *http.Request) httprsp.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(r).Write(w); err != nil {
			panic(err)
		}
	}
}
