package api

import (
	"context"
	"fmt"

	"github.com/google/go-github/v68/github"
)

func FetchWorkflows(ctx context.Context, client *Git) ([]Workflow, error) {
	rawWorkflows, _, err := client.GitClient.Actions.ListWorkflows(ctx, client.owner, client.repo, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting workflows: %v", err)
	}

	var workflows []Workflow
	for _, wf := range rawWorkflows.Workflows {
		workflows = append(workflows, Workflow{
			ID:   wf.GetID(),
			Name: wf.GetName(),
		})
	}

	return workflows, nil
}
func FetchWorkflowRuns(ctx context.Context, client *Git, workflowID int64) ([]Run, error) {
	opts := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: 5},
	}
	runs, _, err := client.GitClient.Actions.ListWorkflowRunsByID(ctx, client.owner, client.repo, workflowID, opts)
	if err != nil {
		return nil, fmt.Errorf("error getting workflow runs: %v", err)
	}

	var workflowRuns []Run
	for _, run := range runs.WorkflowRuns {
		workflowRuns = append(workflowRuns, Run{
			ID:         run.GetID(),
			Status:     run.GetStatus(),
			Conclusion: run.GetConclusion(),
			CreatedAt:  run.GetCreatedAt().Time,
		})
	}

	return workflowRuns, nil
}
