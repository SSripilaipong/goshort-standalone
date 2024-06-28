package httprsp

import "net/http"

var Ok = respondStatusWithHeaderAndBody(http.StatusOK)
