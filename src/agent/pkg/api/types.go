package api

import (
	"fmt"

	"github.com/google/go-github/v68/github"
)

type Workflow struct{
	Name string
	steps []Step
	conclusion string
	status  string
}

type Step struct{
	Name string
	status string
}

type Git struct {
	GitClient *github.Client
	owner string 
	repo string
}


func NewClient(owner, repo, token string) (*Git, error){
	gc := Git{owner: owner, repo:  repo}
	gc.GitClient = github.NewClient(nil).WithAuthToken(token)
	if gc.GitClient == nil {
		fmt.Errorf("Error creating github client")
		return nil, fmt.Errorf("Error creating github client")
	}
	return &gc, nil

}