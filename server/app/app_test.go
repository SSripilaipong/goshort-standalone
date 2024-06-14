package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Join(t *testing.T) {
	t.Run("should join errors", func(t *testing.T) {
		app := newApp([]Applet{
			appletTestError{err: errors.New("yeah")},
			appletTestError{err: errors.New("wee")},
		})
		app.Start()
		errMsg := app.Join().Error()
		assert.Contains(t, errMsg, "yeah")
		assert.Contains(t, errMsg, "wee")
	})

	t.Run("should terminate other app when one completes", func(t *testing.T) {
		app := newApp([]Applet{
			appletTestError{err: errors.New("yeah")},
			&appletTestErrorAfterTerminate{err: errors.New("wee")},
		})
		app.Start()
		errMsg := app.Join().Error()
		assert.Contains(t, errMsg, "yeah")
		assert.Contains(t, errMsg, "wee")
	})
}

type appletTestError struct {
	err error
}

func (a appletTestError) Start() {}

func (a appletTestError) Terminate() {}

func (a appletTestError) Join() error {
	return a.err
}

type appletTestErrorAfterTerminate struct {
	err  error
	wait chan struct{}
}

func (a *appletTestErrorAfterTerminate) Start() {
	a.wait = make(chan struct{}, 1)
}

func (a *appletTestErrorAfterTerminate) Terminate() {
	a.wait <- struct{}{}
}

func (a *appletTestErrorAfterTerminate) Join() error {
	<-a.wait
	return a.err
}
