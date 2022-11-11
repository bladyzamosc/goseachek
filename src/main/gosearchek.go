package main

import "goseachek/src/main/server"

func main() {
	setupServer()
}

func setupServer() {
	server.Server{}.SetupServer()
}
