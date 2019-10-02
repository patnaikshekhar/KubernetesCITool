package actions

import (
	"fmt"

	"github.com/patnaikshekhar/kubernetescitool/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connect(hostname string) (*grpc.ClientConn, error) {
	url, err := config.GetConfig("url")
	if err != nil {
		return nil, err
	}

	if url == "" {
		return nil, fmt.Errorf("Could not find URL in config")
	}

	cert, err := config.GetConfig("cert")
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn

	if cert == "" {
		conn, err = grpc.Dial(url, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(cert, hostname)
		if err != nil {
			return nil, fmt.Errorf("could not load tls cert: %s", err)
		}

		conn, err = grpc.Dial(url, grpc.WithTransportCredentials(creds))
	}

	if err != nil {
		return nil, err
	}

	return conn, nil
}
