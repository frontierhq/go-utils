package azuredevops

import (
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
)

func (a *AzureDevOps) GetPipelineByName(projectName string, name string) (*build.BuildDefinitionReference, error) {
	buildClient, err := build.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getDefinitionsArgs := build.GetDefinitionsArgs{
		Name:    &name,
		Project: &projectName,
	}
	pipelines, err := buildClient.GetDefinitions(a.ctx, getDefinitionsArgs)
	if err != nil {
		return nil, err
	}

	if len(pipelines.Value) == 0 {
		return nil, fmt.Errorf("pipeline with name '%s' not found in project '%s'", name, projectName)
	}
	if len(pipelines.Value) > 1 {
		return nil, fmt.Errorf("multiple pipeline with name '%s' found in project '%s'", name, projectName)
	}

	return &pipelines.Value[0], nil
}
