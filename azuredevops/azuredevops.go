package azuredevops

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type AzureDevOps struct {
	connection *azuredevops.Connection
	ctx        context.Context
}

// NewAzureDevOps creates a new AzureDevOps
func NewAzureDevOps(organisationName string, personalAccessToken string) *AzureDevOps {
	ctx := context.Background()
	organisationUrl := fmt.Sprintf("https://dev.azure.com/%s", organisationName)
	connection := azuredevops.NewPatConnection(organisationUrl, personalAccessToken)
	return &AzureDevOps{
		connection: connection,
		ctx:        ctx,
	}
}

func (a *AzureDevOps) GetPAT() (*string, error) {
	authString := a.connection.AuthorizationString
	if !strings.Contains(authString, "Basic") {
		return nil, fmt.Errorf("authorization string does not contain basic auth")
	}
	bytes, err := base64.StdEncoding.DecodeString(strings.TrimLeft(authString, "Basic "))
	if err != nil {
		return nil, err
	}
	usernamePassword := string(bytes)
	password := strings.Split(usernamePassword, ":")

	return &password[1], nil
}
