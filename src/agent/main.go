package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GHA-Monitor/agent/pkg/api"
	"github.com/GHA-Monitor/agent/pkg/credentials"
)

func main() {
	creds := credentials.Credentials{}
	var err error
	var git *api.Git

	err = creds.Set()
	if err != nil {
		fmt.Printf("Error seting credentials: %s \n", err)
		os.Exit(1)
	}
	git, err = api.NewClient(creds.Owner, creds.Repo, creds.GetToken())
	if err != nil {
		fmt.Println("failed to init github client ")
		os.Exit(1)
	}

	wf, err := api.FetchWorkflows(context.TODO(), git)
	fmt.Printf("printing workflows for %s  repo \n", creds.Repo)
	for _, w := range wf {
		fmt.Println(w.Name)
	}

}
