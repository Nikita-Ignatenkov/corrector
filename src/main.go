package main

import (
	"corrector/app"
	"corrector/check"
)

func main() {
	server := app.NewServer()
	server.Init()
	check.Run(server.DB)
	server.Close()
}
