package main

import (
	"log"
	"os"

	"sprint6/internal/server" // ← УБРАЛ лишний sprint6/
)

func main() {
	logger := log.New(os.Stdout, "morse: ", log.LstdFlags)
	srv := server.New(logger)
	logger.Fatal(srv.Start())
}
