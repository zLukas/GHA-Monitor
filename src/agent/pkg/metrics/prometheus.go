package metrics

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GHA-Monitor/agent/pkg/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
var (
	StepStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gha_step_status",
		Help: "The status of each step in GitHub Actions workflows",
	}, []string{"workflow_name", "run_id", "step_name", "status"})
	reg prometheus.Registerer
)
func InitPrometheus() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(StepStatus)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Println("Starting Prometheus HTTP server on :2112")
	if err := http.ListenAndServe(":2112", nil); err != nil{
		log.Fatalf("Failed to start HTTP server: %v", err)
	}	
}


func ToPrometheusWithSteps(workflows []api.Workflow) {
	for _, wf := range workflows {
		for _, run := range wf.Runs {
			for _, step := range run.Steps {
				status := 0
				if step.Status == "completed" {
					status = 1
				}
				StepStatus.WithLabelValues(wf.Name, strconv.FormatInt(run.ID, 2), step.Name, step.Status).Set(float64(status))
			}
		}
	}
}


