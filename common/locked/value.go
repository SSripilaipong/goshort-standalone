package locked

import "sync"

type Value[T any] struct {
	value T
	lock  sync.RWMutex
}

func NewValue[T any](value T) *Value[T] {
	return &Value[T]{
		value: value,
		lock:  sync.RWMutex{},
	}
}

func (v *Value[T]) Get() T {
	v.lock.RLock()
	defer v.lock.RUnlock()
	return v.value
}

func (v *Value[T]) Set(value T) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = value
}
