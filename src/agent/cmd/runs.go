package cmd

import (
	"context"
	"fmt"

	"github.com/GHA-Monitor/agent/pkg/api"
)

func FetchAllWorkflows(git *api.Git) ([]api.Workflow, error) {
	if git == nil {
		return nil, fmt.Errorf("client is nil")
	}

	workflows, err := api.FetchWorkflows(context.TODO(), git)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch workflows: %w", err)
	}

	return fetchRuns(workflows, context.TODO(), git)
}

func fetchRuns(workflows []api.Workflow, ctx context.Context, client *api.Git) ([]api.Workflow, error) {
	type result struct {
		index int
		runs  []api.Run
		err   error
	}

	runCh := make(chan result, len(workflows))
	for i, workflow := range workflows {
		go func(i int, workflow api.Workflow) {
			runs, err := api.FetchWorkflowRuns(ctx, client, workflow.ID)
			if err != nil {
				runCh <- result{i, nil, fmt.Errorf("failed to fetch runs for workflow %s: %w", workflow.Name, err)}
				return
			}
			for j, run := range runs {
				steps, err := api.FetchRunSteps(ctx, client, run.ID)
				if err != nil {
					runCh <- result{i, nil, fmt.Errorf("failed to fetch steps for run %s: %w", run.ID, err)}
					return
				}
				runs[j].Steps = steps
			}
			runCh <- result{i, runs, nil}
		}(i, workflow)
	}

	for range workflows {
		res := <-runCh
		if res.err != nil {
			return nil, res.err
		}
		workflows[res.index].Runs = res.runs
	}
	return workflows, nil
}
