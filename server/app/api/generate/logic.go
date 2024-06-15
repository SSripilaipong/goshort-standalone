package generate

type Logic interface {
	Generate(url string) (key string, err error)
}

type LogicFunc func(url string) (key string, err error)

func (f LogicFunc) Generate(url string) (key string, err error) {
	return f(url)
}
