package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v68/github"
)

func FetchWorkflows(ctx context.Context, client *Git) ([]Workflow, error) {
	rawWorkflows, resp, err := client.GitClient.Actions.ListWorkflows(ctx, client.owner, client.repo, nil)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: %v ", resp.Status)
	}
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

func FetchRunSteps(ctx context.Context, client *Git, runID int64) ([]Step, error) {
	jobs, _, err := client.GitClient.Actions.ListWorkflowJobs(ctx, client.owner, client.repo, runID, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting workflow jobs: %v", err)
	}

	var steps []Step
	for _, job := range jobs.Jobs {
		for _, step := range job.Steps {
			steps = append(steps, Step{
				Name:   step.GetName(),
				Status: step.GetStatus(),
			})
		}
	}

	return steps, nil
}
