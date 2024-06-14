package wgrp

import "sync"

func WillDone[A, B any](wg *sync.WaitGroup, f func(A) B) func(A) B {
	return func(x A) B {
		defer wg.Done()
		return f(x)
	}
}
