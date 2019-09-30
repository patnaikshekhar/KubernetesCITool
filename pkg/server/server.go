package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"github.com/patnaikshekhar/kubernetescitool/pkg/kube"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"k8s.io/client-go/kubernetes"
)

// Start starts a GRPC KCI Server
func Start(port int, kubeconfig string) {

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("Could not listen %s", err.Error())
	}

	// Connect to Kubernetes
	clientset, err := kube.NewClient(kubeconfig)
	if err != nil {
		log.Fatalf("Could not connect to kubernetes %s", err.Error())
	}

	var grpcServer *grpc.Server

	creds, err := credentials.NewServerTLSFromFile(
		"./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Printf("Warning: Could not load certs")
		grpcServer = grpc.NewServer()
	} else {
		grpcServer = grpc.NewServer(grpc.Creds(creds))
	}

	pb.RegisterKciServer(grpcServer, newKCIServer(clientset))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func newKCIServer(clientset *kubernetes.Clientset) pb.KciServer {
	return kciServer{clientset}
}

type kciServer struct {
	clientset *kubernetes.Clientset
}

func (s kciServer) Build(request *pb.BuildRequest, stream pb.Kci_BuildServer) error {

	podName, err := kube.CreatePod(s.clientset, request)
	if err != nil {
		log.Printf("Error with request : %s", err.Error())
		return err
	}

	err = kube.GetLogs(s.clientset, podName, stream)
	if err != nil {
		log.Printf("Error getting logs: %s", err.Error())
		return err
	}

	return nil
}

func (s kciServer) AddSecret(ctx context.Context,
	request *pb.AddSecretRequest) (*pb.GenericStatus, error) {

	err := kube.CreateSecret(s.clientset, request.Key, request.Value)
	if err != nil {
		log.Printf("Error with request : %s", err.Error())
		return nil, err
	}

	return &pb.GenericStatus{
		Status: "Success",
	}, nil
}
