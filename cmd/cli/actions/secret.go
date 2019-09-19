package actions

import (
	"context"
	"fmt"
	"io/ioutil"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"github.com/patnaikshekhar/kubernetescitool/pkg/config"
	"google.golang.org/grpc"
)

// AddSecret adds a secret to the server
func AddSecret(key string, value string) {

	contents := value

	// Check if value points to a file
	if contents == "" {
		// If File then read contents
		bytes, err := ioutil.ReadFile(key)
		if err != nil {
			fmt.Printf("Could not read file %s", key)
			return
		}
		contents = string(bytes)
	}

	// Invoke service
	request := &pb.AddSecretRequest{
		Key:   key,
		Value: contents,
	}

	err := invokeSecretService(request)
	if err != nil {
		fmt.Printf("Error connecting to KCI Server %s", key)
		return
	}

	fmt.Println("Secret created successfully")
}

func invokeSecretService(request *pb.AddSecretRequest) error {

	url, err := config.GetConfig("url")
	if err != nil {
		return err
	}

	if url == "" {
		return fmt.Errorf("Could not find URL in config")
	}

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewKciClient(conn)

	ctx := context.Background()

	_, err = client.AddSecret(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
