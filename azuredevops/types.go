package azuredevops

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type AzureDevOps struct {
	connection *azuredevops.Connection
	ctx        context.Context
}

type BuildNotFoundError struct {
	name        string
	projectName string
}

func (b *BuildNotFoundError) Error() string {
	return fmt.Sprintf("build definition with name '%s' not found in project '%s'", b.name, b.projectName)
}
