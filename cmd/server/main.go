package main

import "goshort-standalone/server"

func main() {
	if err := server.Start(); err != nil {
		panic(err)
	}
}
