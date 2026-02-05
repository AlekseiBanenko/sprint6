package main

import (
	"log"
	"os"

	"server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	srv := server.New(logger)
	logger.Fatal("Server failed: ", srv.Start())
}
