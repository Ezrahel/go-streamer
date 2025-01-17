package main

import (
	"log"

	"github.com/ezrahel/go-streamer/internals/handlers/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error)
	}
}
