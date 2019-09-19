package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"github.com/patnaikshekhar/kubernetescitool/pkg/kube"
	"google.golang.org/grpc"
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

	grpcServer := grpc.NewServer()
	pb.RegisterKciServer(grpcServer, newKCIServer(clientset))

	grpcServer.Serve(lis)
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
