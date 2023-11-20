package azuredevops

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type AzureDevOps struct {
	connection *azuredevops.Connection
	ctx        context.Context
}

type CustomQueueBuildArgs struct {
	Definition         CustomDefinition  `json:"definition"`
	SourceBranch       string            `json:"sourceBranch"`
	TemplateParameters map[string]string `json:"templateParameters"`
}

type CustomDefinition struct {
	ID *int `json:"id"`
}
