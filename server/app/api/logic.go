package api

import (
	"goshort-standalone/server/app/api/access"
	"goshort-standalone/server/app/api/generate"
)

type logic interface {
	generate.Logic
	access.Logic
}
