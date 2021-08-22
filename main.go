package main

import (
	"os"
	"twit/server"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Start(port)
}
