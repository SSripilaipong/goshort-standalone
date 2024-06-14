package server

import (
	"goshort-standalone/server/app"
	"goshort-standalone/server/app/api"
	"goshort-standalone/server/app/gateway"
	"goshort-standalone/server/app/membermgr"
)

func Start() error {
	return app.New(
		[]app.Applet{
			api.New(),
			gateway.New(),
			membermgr.New(),
		},
	).Join()
}
