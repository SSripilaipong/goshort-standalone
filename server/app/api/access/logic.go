package access

type Logic interface {
	Access(key string) (url string, err error)
}

type LogicFunc func(key string) (url string, err error)

func (f LogicFunc) Access(key string) (url string, err error) {
	return f(key)
}
