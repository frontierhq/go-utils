package azuredevops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	azuredevops "github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
)

// GetBuild gets a Build
func (a *AzureDevOps) GetBuild(projectName string, buildId *int) (*build.Build, error) {
	client, err := build.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getBuildArgs := build.GetBuildArgs{
		BuildId: buildId,
		Project: &projectName,
	}
	return client.GetBuild(a.ctx, getBuildArgs)
}

// GetBuildDefinitionByName gets a BuildDefinitionReference by name
func (a *AzureDevOps) GetBuildDefinitionByName(projectName string, name string) (*build.BuildDefinitionReference, error) {
	client, err := build.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getDefinitionsArgs := build.GetDefinitionsArgs{
		Name:    &name,
		Project: &projectName,
	}
	definitions, err := client.GetDefinitions(a.ctx, getDefinitionsArgs)
	if err != nil {
		return nil, err
	}

	if len(definitions.Value) == 0 {
		return nil, fmt.Errorf("build definition with name '%s' not found in project '%s'", name, projectName)
	}
	if len(definitions.Value) > 1 {
		return nil, fmt.Errorf("multiple build definitions with name '%s' found in project '%s'", name, projectName)
	}

	return &definitions.Value[0], nil
}

type CustomQueueBuildArgs struct {
	Definition         CustomDefinition  `json:"definition"`
	SourceBranch       string            `json:"sourceBranch"`
	TemplateParameters map[string]string `json:"templateParameters"`
}

type CustomDefinition struct {
	ID *int `json:"id"`
}

// QueueBuild queues and returns a new Build
func (a *AzureDevOps) QueueBuild(projectName string, definitionId *int, sourceBranch string, templateParameters map[string]string, tags []string) (*build.Build, error) {
	buildClient, err := build.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	// queueBuildArgs := build.QueueBuildArgs{
	// 	Build: &build.Build{
	// 		Definition: &build.DefinitionReference{
	// 			Id: definitionId,
	// 		},
	// 		SourceBranch:       &sourceBranch,
	// 		TemplateParameters: &templateParameters,
	// 	},
	// 	Project: &projectName,
	// }
	// queuedBuild, err := client.QueueBuild(a.ctx, queueBuildArgs)

	// The Build type (at QueueBuildArgs.Build) doesn't include TemplateParameters,
	// so we can't use this function to trigger pipelines that take params.
	// In the meantime, this workaround will do:

	client := azuredevops.NewClient(a.connection, a.connection.BaseUrl)

	queueBuildArgs := &CustomQueueBuildArgs{
		Definition: CustomDefinition{
			ID: definitionId,
		},
		SourceBranch:       sourceBranch,
		TemplateParameters: templateParameters,
	}

	body, err := json.Marshal(queueBuildArgs)
	if err != nil {
		return nil, err
	}

	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")

	routeValues := map[string]string{
		"project": projectName,
	}

	queryParams := url.Values{}

	response, err := client.Send(a.ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		if _, ok := err.(azuredevops.WrappedError); ok {
			wrappedError := err.(azuredevops.WrappedError)
			if *wrappedError.TypeKey == "BuildRequestValidationFailedException" {
				validationResults := (*wrappedError.CustomProperties)["ValidationResults"]
				builder := &strings.Builder{}
				for i, v := range validationResults.([]interface{}) {
					message := v.(map[string]interface{})["message"]
					index := i + 1
					builder.WriteString(fmt.Sprintf("\n%d) %s", index, message))
				}
				return nil, fmt.Errorf("%s %s", err.Error(), builder.String())
			}
		}
		return nil, err
	}

	var queuedBuild build.Build
	err = client.UnmarshalBody(response, &queuedBuild)
	if err != nil {
		return nil, err
	}

	addBuildTagsArgs := build.AddBuildTagsArgs{
		BuildId: queuedBuild.Id,
		Project: &projectName,
		Tags:    &tags,
	}
	_, err = buildClient.AddBuildTags(a.ctx, addBuildTagsArgs)
	if err != nil {
		return nil, err
	}

	return &queuedBuild, err
}

// WaitForBuild waits for a Build to complete
func (a *AzureDevOps) WaitForBuild(projectName string, buildId *int, attempts uint, interval int) error {
	err := retry.Do(
		func() error {
			var err error
			build, err := a.GetBuild(projectName, buildId)
			if err != nil {
				if _, ok := err.(azuredevops.WrappedError); ok {
					wrappedError := err.(azuredevops.WrappedError)
					if *wrappedError.TypeKey == "BuildNotFoundException" {
						return retry.Unrecoverable(err)
					}
				}
				return err
			}

			switch string(*build.Status) {
			case "completed":
				return nil
			case "postponed":
				return retry.Unrecoverable(fmt.Errorf("build '%s' has been postponed", *build.BuildNumber))
			default:
				return fmt.Errorf("build '%s' has status '%s'", *build.BuildNumber, *build.Status)
			}
		},
		retry.Attempts(attempts),
		retry.Delay(time.Duration(interval)*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	)

	return err
}
