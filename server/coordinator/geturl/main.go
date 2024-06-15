package geturl

import "fmt"

func New() func(key string) (url string, err error) {
	return func(key string) (url string, err error) {
		return fmt.Sprintf("https://www.google.com/search?q=%s", key), nil
	}
}
