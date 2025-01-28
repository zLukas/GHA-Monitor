package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GHA-Monitor/agent/cmd"
	"github.com/GHA-Monitor/agent/pkg/api"
)

func main() {
	var err error
	var git *api.Git
	git, err = cmd.GitHubInit()
	if err != nil {
		fmt.Printf("Error seting credentials: %s \n", err)
		os.Exit(1)
	}
	wf, err := api.FetchWorkflows(context.TODO(), git)
	if err != nil {
		fmt.Printf("Error fetching workflows: %s \n", err)
		os.Exit(1)
	}
	for _, w := range wf {
		fmt.Println(w.Name)
	}

}
