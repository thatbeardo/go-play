package main

import (
	"log"
	"os"

	"github.com/thatbeardo/go-play/server"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.Flags())
	server.StartServer(logger)
}
