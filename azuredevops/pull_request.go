package azuredevops

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

// AbandonPullRequest abandons a GitPullRequest
func (a *AzureDevOps) AbandonPullRequest(projectName string, repositoryName string, pullRequestId int) (*git.GitPullRequest, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	pullRequest := git.GitPullRequest{
		Status: &git.PullRequestStatusValues.Abandoned,
	}

	updatePullRequestArgs := git.UpdatePullRequestArgs{
		GitPullRequestToUpdate: &pullRequest,
		Project:                &projectName,
		PullRequestId:          &pullRequestId,
		RepositoryId:           &repositoryName,
	}
	return client.UpdatePullRequest(a.ctx, updatePullRequestArgs)
}

// FindPullRequest finds a GitPullRequest
func (a *AzureDevOps) FindPullRequest(projectName string, repositoryName string, sourceRefName string, targetRefName string) (*git.GitPullRequest, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	repository, err := a.GetRepository(projectName, repositoryName)
	if err != nil {
		return nil, err
	}

	getPullRequestsArgs := git.GetPullRequestsArgs{
		Project:      &projectName,
		RepositoryId: &repositoryName,
		SearchCriteria: &git.GitPullRequestSearchCriteria{
			RepositoryId:  repository.Id,
			SourceRefName: &sourceRefName,
			TargetRefName: &targetRefName,
		},
	}

	pullRequests, err := client.GetPullRequests(a.ctx, getPullRequestsArgs)
	if err != nil {
		return nil, err
	}

	return &(*pullRequests)[0], nil
}

// CreatePullRequest creates a GitPullRequest
func (a *AzureDevOps) CreatePullRequest(projectName string, repositoryName string, sourceRefName string, targetRefName string, title string) (*git.GitPullRequest, error) {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	pullRequest := git.GitPullRequest{
		SourceRefName: &sourceRefName,
		TargetRefName: &targetRefName,
		Title:         &title,
	}

	createPullRequestArgs := git.CreatePullRequestArgs{
		GitPullRequestToCreate: &pullRequest,
		Project:                &projectName,
		RepositoryId:           &repositoryName,
	}
	return client.CreatePullRequest(a.ctx, createPullRequestArgs)
}
