package azuredevops

import (
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"golang.org/x/exp/slices"
)

// GetRepositories gets all GitRepository
func (a *AzureDevOps) GetRepositories(projectName string) (*[]git.GitRepository, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getRepositoriesArgs := git.GetRepositoriesArgs{
		Project: &projectName,
	}
	return client.GetRepositories(a.ctx, getRepositoriesArgs)
}

// GetRepository gets a GitRepository
func (a *AzureDevOps) GetRepository(projectName string, name string) (*git.GitRepository, error) {
	repositories, err := a.GetRepositories(projectName)
	if err != nil {
		return nil, err
	}

	findRepositoryFunc := func(r git.GitRepository) bool { return *r.Name == name }
	repositoryIdx := slices.IndexFunc(*repositories, findRepositoryFunc)

	if repositoryIdx == -1 {
		return nil, fmt.Errorf("repository with name '%s' not found in project '%s'", name, projectName)
	}

	return &(*repositories)[repositoryIdx], nil
}
