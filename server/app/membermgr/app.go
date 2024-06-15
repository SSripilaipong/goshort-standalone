package membermgr

import "goshort-standalone/common/signal"

type App struct {
	termination signal.Once
}

func (a App) Start() {
}

func (a App) Terminate() {
	a.termination.Happens()
}

func (a App) Join() error {
	a.termination.Wait()
	return nil
}
