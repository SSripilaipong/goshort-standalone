package access

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"goshort-standalone/common/httprsp"
	"goshort-standalone/common/httptst"
	"goshort-standalone/common/rslt"
)

func TestNewHandler(t *testing.T) {
	t.Run("should pass key to access logic", func(t *testing.T) {
		var capturedKey string

		httptst.MakeRequest(httptst.HandlerServer("/abc/{key}", handlerFuncForTest(func(key string) rslt.Of[string] {
			capturedKey = key
			return rslt.Of[string]{}
		})), httptest.NewRequest(http.MethodGet, "/abc/my-key", nil))

		assert.Equal(t, capturedKey, "my-key")
	})

	t.Run("should respond from access logic", func(t *testing.T) {
		response := httptst.MakeRequest(httptst.HandlerServer("/abc/{key}", handlerFuncForTest(func(key string) rslt.Of[string] {
			return rslt.Value[string]("my-url")
		})), httptest.NewRequest(http.MethodGet, "/abc/my-key", nil))

		assert.Equal(t, http.StatusPermanentRedirect, response.Code)
		assert.Equal(t, []string{"my-url"}, response.Header()["Location"])
	})
}

func handlerFuncForTest(f LogicFunc) http.HandlerFunc {
	return httprsp.HandlerFunc(NewHandler(f))
}
