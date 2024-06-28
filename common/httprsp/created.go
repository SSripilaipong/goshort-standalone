package httprsp

import "net/http"

var Created = respondStatusWithHeaderAndBody(http.StatusCreated)
