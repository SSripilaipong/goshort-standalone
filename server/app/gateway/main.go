package gateway

import "goshort-standalone/common/signal"

func New() App {
	return App{
		termination: signal.NewOnce(),
	}
}
