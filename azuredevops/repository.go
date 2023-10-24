package azuredevops

import (
	"fmt"
	"os"
	"path/filepath"

	egit "github.com/frontierdigital/utils/git/external_git"
	"github.com/frontierdigital/utils/output"
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

// configureRepository configures a repository
func configureRepository(repo *egit.ExternalGit, gitEmail string, gitUsername string) error {
	err := repo.SetConfig("user.email", gitEmail)
	if err != nil {
		return err
	}
	err = repo.SetConfig("user.name", gitUsername)
	if err != nil {
		return err
	}
	return nil
}

// createRepositoryIfNotExists creates a repository if it does not exist
func createRepositoryIfNotExists(a *AzureDevOps, projectName string, repoName string, gitEmail string, gitUsername string, adoPat string) (*git.GitRepository, *string, error) {
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
		repoPath, err := os.MkdirTemp("", "")
		if err != nil {
			return nil, nil, err
		}
		repo := egit.NewGit(repoPath)
		err = repo.CloneOverHttp(*r.RemoteUrl, adoPat, "x-oauth-basic")
		if err != nil {
			return nil, nil, err
		}
		err = configureRepository(repo, gitEmail, gitUsername)
		if err != nil {
			return nil, nil, err
		}
		return r, &repoPath, nil
	}

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

	localPath, err := initRepository("frontierdigital", projectName, repoName, gitEmail, gitUsername, adoPat)

	if err != nil {
		return nil, nil, err
	}

	return r, localPath, nil
}

// initRepository creates a main branch
func initRepository(organisationName string, projectName string, repoName string, gitEmail string, gitUsername string, adoPat string) (*string, error) {
	repoUrl := fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s", organisationName, projectName, repoName)

	repoPath, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}
	repo := egit.NewGit(repoPath)
	err = repo.CloneOverHttp(repoUrl, adoPat, "x-oauth-basic")
	if err != nil {
		return nil, err
	}
	err = configureRepository(repo, gitEmail, gitUsername)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(filepath.Join(repoPath, "README.md"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = repo.AddAll()
	if err != nil {
		return nil, err
	}

	commitMessage := "Initial Commit"
	_, err = repo.Commit(commitMessage)
	if err != nil {
		return nil, err
	}

	err = repo.Push(false)
	if err != nil {
		return nil, err
	}

	output.PrintlnfInfo("Pushed.")

	return &repoPath, nil
}
