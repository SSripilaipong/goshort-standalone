package api

type App struct{}

func (a App) Start() {
}

func (a App) Terminate() {
}

func (a App) Join() error {
	return nil
}
