package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"google.golang.org/grpc"
)

// Start starts a GRPC KCI Server
func Start(port int) {

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("Could not listen %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKciServer(grpcServer, newKCIServer())

	grpcServer.Serve(lis)
}

func newKCIServer() pb.KciServer {
	return kciServer{}
}

type kciServer struct{}

func (s kciServer) Build(ctx context.Context, request *pb.BuildRequest) (
	*pb.BuildResponse, error) {

	log.Printf("Received BuildRequest %+v", request)

	// Get request

	// Connect to Kubernetes

	// Start Pod

	return &pb.BuildResponse{Update: "Done"}, nil
}
