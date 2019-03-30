package main

import (
	"fmt"
	cmd "github.com/sunkuet02/go-grpc-rest/pkg/cmd/server"
	"os"
)

func main() {
	err := cmd.RunServer()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1);
	}
}