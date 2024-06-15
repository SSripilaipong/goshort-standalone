package getkeyorcreate

import "fmt"

func New() func(url string) (key string, err error) {
	return func(url string) (key string, err error) {
		return fmt.Sprintf("hash(%s)", url), nil
	}
}
