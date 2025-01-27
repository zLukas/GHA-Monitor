package api

import (
	"fmt"
	"time"

	"github.com/google/go-github/v68/github"
)

type Git struct {
	GitClient *github.Client
	owner     string
	repo      string
}

type Workflow struct {
	ID   int64
	Name string
}

type Run struct {
	ID         int64
	Status     string
	Conclusion string
	CreatedAt  time.Time
}

func NewClient(owner, repo, token string) (*Git, error) {
	gc := Git{owner: owner, repo: repo}
	gc.GitClient = github.NewClient(nil).WithAuthToken(token)
	if gc.GitClient == nil {
		fmt.Errorf("Error creating github client")
		return nil, fmt.Errorf("Error creating github client")
	}
	return &gc, nil

}
