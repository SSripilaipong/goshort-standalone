package app

import (
	"errors"
	"sync"

	"goshort-standalone/common/chn"
	"goshort-standalone/common/fn"
	"goshort-standalone/common/slc"
	"goshort-standalone/common/wgrp"
)

type App struct {
	applets []Applet
}

func newApp(applets []Applet) App {
	return App{applets: applets}
}

func (a App) Start() {
	slc.Apply(startApplet, a.applets)
}

func (a App) Join() error {
	return errors.Join(chn.All(a.joinAndTerminateAll())...)
}

func (a App) joinAndTerminateAll() <-chan error {
	var wg sync.WaitGroup
	wg.Add(len(a.applets))

	errCh := make(chan error, len(a.applets))
	joinOneAndTerminate := fn.Unvoid[any](a.joinOneAndTerminateAll(errCh))

	go func() {
		defer close(errCh)
		wg.Wait()
	}()

	slc.Apply(fn.GoFn(wgrp.WillDone(&wg, joinOneAndTerminate)), a.applets)
	return errCh
}

func (a App) joinOneAndTerminateAll(errCh chan<- error) func(p Applet) {
	return func(p Applet) {
		err := p.Join()
		if err != nil {
			errCh <- err
		}
		a.terminateAll()
	}
}

func (a App) terminateAll() {
	slc.Apply(terminateApplet, a.applets)
}
