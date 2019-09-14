package main

import (
	"flag"

	"github.com/patnaikshekhar/kubernetescitool/pkg/server"
)

func main() {
	port := flag.Int("port", 10000, "The server port")

	flag.Parse()

	// Start GRPC Server
	server.Start(*port)
}
