package api

import (
	"net/http"

	"goshort-standalone/common/locked"
	"goshort-standalone/common/signal"
)

type App struct {
	server *http.Server
	done   signal.Once
	err    *locked.Value[error]
}

func newApp(port uint16, handler http.Handler) *App {
	return &App{
		server: newHttpServer(port, handler),
		done:   signal.NewOnce(),
		err:    locked.NewValue[error](nil),
	}
}

func (a *App) Start() {
	go func() {
		defer a.done.Happens()
		err := a.server.ListenAndServe()
		if err != nil {
			a.err.Set(err)
		}
	}()
}

func (a *App) Terminate() {
	_ = a.server.Close()
}

func (a *App) Join() error {
	a.done.Wait()
	return a.err.Get()
}
