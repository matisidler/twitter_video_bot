package main

import (
	"os"
	"twit/funcs"
	"twit/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	funcs.Testing()
	server.Start(port)

}
