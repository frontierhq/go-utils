package azuredevops

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

// GetFileContent gets content in a file over API.
func (a *AzureDevOps) GetFileContent(projectName string, repoName string, version string) (*git.GitItem, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	path := "README.md"
	includeContent := true
	gitVersionDescriptor := git.GitVersionDescriptor{
		VersionType: &git.GitVersionTypeValues.Tag,
		Version:     &version,
	}
	getItemArgs := git.GetItemArgs{
		RepositoryId:      &repoName,
		Project:           &projectName,
		Path:              &path,
		IncludeContent:    &includeContent,
		VersionDescriptor: &gitVersionDescriptor,
	}
	return client.GetItem(a.ctx, getItemArgs)
}
