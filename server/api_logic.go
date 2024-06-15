package server

import (
	"goshort-standalone/server/app/api/access"
	"goshort-standalone/server/app/api/generate"
	"goshort-standalone/server/coordinator/getkeyorcreate"
	"goshort-standalone/server/coordinator/geturl"
)

func newApiLogic() apiLogic {
	return apiLogic{
		generateLogic: getkeyorcreate.New(),
		accessLogic:   geturl.New(),
	}
}

type apiLogic struct {
	generateLogic
	accessLogic
}
type generateLogic = generate.LogicFunc
type accessLogic = access.LogicFunc
