package api

import (
	"context"
	"fmt"
)

func Fetch_workflow_metadatachan(ctx context.Context, client *Git) ([]Workflow, error){
	rawWorkflows, _, err:= client.GitClient.Actions.ListWorkflows(ctx, client.owner, client.repo, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting workflows: %v", err)
	}
	return nil, nil
}