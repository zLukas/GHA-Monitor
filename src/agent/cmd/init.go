package cmd

import (
	"fmt"

	"github.com/GHA-Monitor/agent/pkg/api"
	"github.com/GHA-Monitor/agent/pkg/credentials"
)

func GitHubInit() (*api.Git, error) {

	creds := credentials.Credentials{}
	var err error
	var git *api.Git

	err = creds.Set()
	if err != nil {
		fmt.Printf("Error seting credentials: %s \n", err)
		return nil, err
	}
	git, err = api.NewClient(creds.Owner, creds.Repo, creds.GetToken())
	if err != nil {
		fmt.Println("failed to init github client ")
		return nil, err
	}
	return git, nil
}
