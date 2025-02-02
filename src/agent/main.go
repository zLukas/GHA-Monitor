package main

import (
	"encoding/json"
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
	fmt.Println("Agent is running...")
	//cmd.Serve(git)
	wfs, err := cmd.FetchAllWorkflows(git)
	if err != nil {
		fmt.Printf("Error fetching workflow content: %s \n", err)
		os.Exit(1)
	}
	file, err := os.Create("workflows.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %s \n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(wfs); err != nil {
		fmt.Printf("Error encoding JSON: %s \n", err)
	}
}
