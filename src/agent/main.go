package main

import (
	"fmt"
	"os"

	"github.com/GHA-Monitor/agent/pkg/api"
	"github.com/GHA-Monitor/agent/pkg/credentials"
)

func main() {
	creds := credentials.Credentials{}
	var err error

	err = creds.Set()
	if err != nil {
		fmt.Printf("Error seting credentials: %s \n", err)
		os.Exit(1)
	}
	 _ ,err = api.NewClient(creds.Owner, creds.Repo, creds.GetToken())
	if err != nil {
		fmt.Println("failed to init github client ")
		os.Exit(1)
	}
}