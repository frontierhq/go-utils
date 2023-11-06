package azuredevops

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// getProjectUUID gets the UUID of a project
func (a *AzureDevOps) getProjectUUID(projectName string) (*uuid.UUID, error) {
	client, err := core.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getProjectsArgs := core.GetProjectsArgs{}
	projects, err := client.GetProjects(a.ctx, getProjectsArgs)
	if err != nil {
		return nil, err
	}

	for _, project := range projects.Value {
		if *project.Name == projectName {
			return project.Id, nil
		}
	}

	return nil, fmt.Errorf("project %s not found", projectName)
}
