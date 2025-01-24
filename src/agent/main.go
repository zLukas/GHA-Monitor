package main

import (
	"fmt"
	"os"

	"github.com/GHA-Monitor/agent/pkg/credentials"
)

func main() {
	creds := credentials.Credentials{}
	err := creds.Set()
	if err != nil {
		fmt.Printf("Error seting credentials: %s \n", err)
		os.Exit(1)
	}
	fmt.Printf("token %s\n", creds.GetToken())
}