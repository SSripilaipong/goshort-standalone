package generate

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"goshort-standalone/common/fn"
	"goshort-standalone/common/httprsp"
	"goshort-standalone/common/httptst"
	"goshort-standalone/common/rslt"
)

func TestNewHandler(t *testing.T) {
	t.Run("should pass url to generate logic", func(t *testing.T) {
		var capturedUrl string

		httptst.MakeRequest(httptst.HandlerServer("/gen", handlerFuncForTest(func(url string) rslt.Of[string] {
			capturedUrl = url
			return rslt.Of[string]{}
		})), httptest.NewRequest(http.MethodPost, "/gen", strings.NewReader(`{"url":"google.com"}`)))

		assert.Equal(t, capturedUrl, "google.com")
	})

	t.Run("should respond from generate logic", func(t *testing.T) {
		response := testRequestWithBody(
			fn.Const[string](rslt.Value[string]("abc123")),
			strings.NewReader(`{"url":"yahoo.com"}`),
		)

		body := httptst.JsonResponse(response.Body)
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, "abc123", body["key"])
		assert.Equal(t, "yahoo.com", body["url"])
	})
}

func testRequestWithBody(f LogicFunc, body io.Reader) *httptest.ResponseRecorder {
	return httptst.MakeRequest(
		httptst.HandlerServer("/gen", handlerFuncForTest(f)),
		httptest.NewRequest(http.MethodPost, "/gen", body),
	)
}

func handlerFuncForTest(f LogicFunc) http.HandlerFunc {
	return httprsp.HandlerFunc(NewHandler(f))
}
