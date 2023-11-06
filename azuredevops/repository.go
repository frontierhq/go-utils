package azuredevops

import (
	"fmt"
	"os"

	egit "github.com/frontierdigital/utils/git/external_git"
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

// createRepositoryIfNotExists creates a repository if it does not exist
func (a *AzureDevOps) createRepositoryIfNotExists(projectName string, repoName string, gitEmail string, gitUsername string) (*git.GitRepository, *string, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, nil, err
	}

	getRepositoryArgs := git.GetRepositoryArgs{
		RepositoryId: &repoName,
		Project:      &projectName,
	}

	r, err := client.GetRepository(a.ctx, getRepositoryArgs)

	if err == nil {
		repoPath, err := a.checkoutAndConfigure(*r.RemoteUrl, gitEmail, gitUsername)
		if err != nil {
			return nil, nil, err
		}

		return r, repoPath, nil
	}

	// TODO: Check that err is a GitRepositoryNotFound error

	createRepositoryArgs := git.CreateRepositoryArgs{
		GitRepositoryToCreate: &git.GitRepositoryCreateOptions{
			Name: &repoName,
		},
		Project: &projectName,
	}

	r, err = client.CreateRepository(a.ctx, createRepositoryArgs)
	if err != nil {
		return nil, nil, err
	}

	repoPath, err := a.initRepository(*r.RemoteUrl, gitEmail, gitUsername)
	if err != nil {
		return nil, nil, err
	}

	return r, repoPath, nil
}

func (a *AzureDevOps) checkoutAndConfigure(remoteUrl string, gitEmail string, gitUsername string) (*string, error) {
	repoPath, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	repo := egit.NewGit(repoPath)
	pat, err := a.GetPAT()
	if err != nil {
		return nil, err
	}

	err = repo.CloneOverHttp(remoteUrl, *pat, "x-oauth-basic")
	if err != nil {
		return nil, err
	}

	err = repo.Configure(gitEmail, gitUsername)
	if err != nil {
		return nil, err
	}

	repoPath = repo.GetRepositoryPath()

	return &repoPath, nil
}

// initRepository creates a main branch
func (a *AzureDevOps) initRepository(remoteUrl string, gitEmail string, gitUsername string) (*string, error) {
	repoPath, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	repo := egit.NewGit(repoPath)

	pat, err := a.GetPAT()
	if err != nil {
		return nil, err
	}

	err = repo.CloneOverHttp(remoteUrl, *pat, "x-oauth-basic")
	if err != nil {
		return nil, err
	}

	err = repo.Configure(gitEmail, gitUsername)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(repo.GetFilePath("README.md"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = repo.AddAll()
	if err != nil {
		return nil, err
	}

	commitMessage := "Initial commit."
	_, err = repo.Commit(commitMessage)
	if err != nil {
		return nil, err
	}

	err = repo.Push(false)
	if err != nil {
		return nil, err
	}

	return &repoPath, nil
}
