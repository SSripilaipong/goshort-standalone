package httprsp

import "net/http"

func HandlerFunc(f func(request *http.Request) Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(r).Write(w); err != nil {
			panic(err)
		}
	}
}
