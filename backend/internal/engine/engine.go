// the idea is to have the engine here that will execute the commands
// and return the result to the client
package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type PipelineConfig struct {
	Pipeline Pipeline `yaml:"pipeline"`
}

type Pipeline struct {
	Name   string  `yaml:"name"`
	Stages []Stage `yaml:"stages"`
}

type Stage struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name   string `yaml:"name"`
	Action string `yaml:"action"`
}

func main() {
	fmt.Println("Starting pipeline...")
	pipelinePath := "pipeline.yml"

	pipelineData, err := os.ReadFile(pipelinePath)
	if err != nil {
		fmt.Printf("Error reading pipeline file: %v\n", err)
		return
	}

	var config PipelineConfig
	err = yaml.Unmarshal(pipelineData, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling pipeline data: %v\n", err)
		return
	}

	pipeline := config.Pipeline
	fmt.Printf("Processing pipeline: %s\n", pipeline.Name)

	for _, stage := range pipeline.Stages {
		fmt.Printf("Processing stage: %s\n", stage.Name)
		for _, step := range stage.Steps {
			executeStep(step)
		}
	}

	fmt.Println("Pipeline execution completed.")
}

func executeStep(step Step) {
	fmt.Printf("Executing step: %s, Action: %s\n", step.Name, step.Action)
	switch step.Action {
	case "git clone https://github.com/myorg/myapp.git":
		fmt.Printf("Cloning repository for step: %s\n", step.Name)
	case "npm install":
		fmt.Printf("Installing dependencies for step: %s\n", step.Name)
	case "npm run build":
		fmt.Printf("Building code for step: %s\n", step.Name)
	default:
		fmt.Printf("Unknown action for step: %s\n", step.Name)
	}
}
