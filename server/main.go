package server

import (
	"goshort-standalone/server/app"
	"goshort-standalone/server/app/api"
	"goshort-standalone/server/app/gateway"
	"goshort-standalone/server/app/interrupt"
	"goshort-standalone/server/app/membermgr"
)

func Start() error {
	a := app.New(
		[]app.Applet{
			api.New(newApiLogic()),
			interrupt.New(),
			gateway.New(),
			membermgr.New(),
		},
	)
	a.Start()
	return a.Join()
}
