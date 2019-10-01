package actions

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"gopkg.in/yaml.v2"
)

// Build calls the gRPC interface on the controlplane to start the build
// process
func Build(filename string) {
	// Read the file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading input file %s", err.Error())
	}

	// Convert Contents to format
	var request pb.BuildRequest
	err = yaml.UnmarshalStrict(contents, &request)

	if err != nil {
		log.Fatalf("Could not parse file %s", err.Error())
	}

	if err := StartBuild(&request); err != nil {
		log.Fatalf("Error invoking server %s", err.Error())
	}

}

// StartBuild starts a build by calling the gRPC interface
func StartBuild(request *pb.BuildRequest) error {

	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewKciClient(conn)

	ctx := context.Background()

	responseStream, err := client.Build(ctx, request)
	if err != nil {
		return err
	}

	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			fmt.Println("Build Completed")
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf("Running Step %d\n\n", response.Step)
		fmt.Println(response.Data)
	}
}
