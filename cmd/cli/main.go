package main

import (
	"flag"
	"fmt"

	"github.com/patnaikshekhar/kubernetescitool/cmd/cli/actions"
	"github.com/patnaikshekhar/kubernetescitool/pkg/config"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if args[0] == "config" {
		if len(args) >= 3 {
			config.AddConfig(args[1], args[2])
		} else {
			fmt.Println("You must pass a key and value with config")
		}
	} else if args[0] == "secret" {
		if len(args) >= 2 {
			if len(args) >= 3 {
				actions.AddSecret(args[1], args[2])
			} else {
				actions.AddSecret(args[1], "")
			}
		} else {
			fmt.Println(
				"You must either provide a file name or a key-value pair")
		}
	} else {
		// Example kci build.yaml
		if len(args) >= 1 {
			actions.Build(args[0])
		} else {
			fmt.Println("You must pass the build filename")
		}
	}
}
