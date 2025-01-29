package cmd

import (
	"fmt"
	"time"

	"github.com/GHA-Monitor/agent/pkg/api"
	"github.com/GHA-Monitor/agent/pkg/metrics"
)

func Serve(git *api.Git) {
	metrics.InitPrometheus()
	for {
		
		wfs, err := FetchAllWorkflows(git)
		if err != nil {
			fmt.Printf("Error fetching workflows: %s \n", err)
		}
		metrics.ToPrometheusWithSteps(wfs)
		time.Sleep(5 * time.Second)
	}
}