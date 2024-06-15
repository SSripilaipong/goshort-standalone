package signal

type Once chan struct{}

func NewOnce() Once {
	return make(Once)
}

func (q Once) Happens() {
	defer func() {
		_ = recover()
	}()
	close(q)
}

func (q Once) Wait() {
	<-q
}

func (q Once) UntilHappens() <-chan struct{} {
	return q
}
