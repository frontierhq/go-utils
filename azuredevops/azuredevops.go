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

func NewAzureDevOps(organisationName string, personalAccessToken string) *AzureDevOps {
	ctx := context.Background()
	organisationUrl := fmt.Sprintf("https://dev.azure.com/%s", organisationName)
	connection := azuredevops.NewPatConnection(organisationUrl, personalAccessToken)
	return &AzureDevOps{
		connection: connection,
		ctx:        ctx,
	}
}
