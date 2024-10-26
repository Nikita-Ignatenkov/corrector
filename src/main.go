package main

import (
	"corrector/app"
	"corrector/check"
	"corrector/cmd"
	"log"
)

func main() {
	server := app.NewServer()
	server.Init()
	check.Run(server.DB)
	//parser.ParseCategory(687)
	if err := cmd.Root(server).Execute(); err != nil {
		log.Fatal(err)
	}
	server.Close()
}
