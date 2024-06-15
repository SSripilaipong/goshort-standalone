package interrupt

import (
	"errors"
	"os"
	osSignal "os/signal"

	"goshort-standalone/common/signal"
)

type App struct {
	termination signal.Once
}

func newApp() App {
	return App{
		termination: signal.NewOnce(),
	}
}

func (a App) Start() {
}

func (a App) Terminate() {
	a.termination.Happens()
}

func (a App) Join() error {
	c := make(chan os.Signal, 1)
	osSignal.Notify(c, os.Interrupt)
	interrupt := make(chan struct{})
	go func() {
		for range c {
			close(interrupt)
		}
	}()

	select {
	case <-interrupt:
		return errors.New("os interrupted")
	case <-a.termination.UntilHappens():
	}
	return nil
}
