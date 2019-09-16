package main

import (
	"flag"
	"os"

	"github.com/patnaikshekhar/kubernetescitool/pkg/server"
)

func main() {
	port := flag.Int("port", 10000, "The server port")
	kubeconfig := flag.String("kubeconfig", "", "Location of kubeconfig")

	flag.Parse()

	if *kubeconfig == "" {
		kubeconfigFromEnv := os.Getenv("KUBECONFIG")
		kubeconfig = &kubeconfigFromEnv
	}

	// Start GRPC Server
	server.Start(*port, *kubeconfig)
}
