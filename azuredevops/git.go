package azuredevops

import (
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

// GetFileContent gets content in a file over API.
func (a *AzureDevOps) GetFileContent(projectName string, repoName string, version string, filepath string, versionType string) (*git.GitItem, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	var vt git.GitVersionType
	switch versionType {
	case "branch":
		vt = git.GitVersionTypeValues.Branch
	case "commit":
		vt = git.GitVersionTypeValues.Commit
	case "tag":
		vt = git.GitVersionTypeValues.Tag
	default:
		return nil, fmt.Errorf("unknown version type: %s", versionType)
	}

	includeContent := true
	gitVersionDescriptor := git.GitVersionDescriptor{
		VersionType: &vt,
		Version:     &version,
	}
	getItemArgs := git.GetItemArgs{
		RepositoryId:      &repoName,
		Project:           &projectName,
		Path:              &filepath,
		IncludeContent:    &includeContent,
		VersionDescriptor: &gitVersionDescriptor,
	}
	return client.GetItem(a.ctx, getItemArgs)
}
