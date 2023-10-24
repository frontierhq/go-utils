package azuredevops

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
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

// CompletePullRequest completes a GitPullRequest
func (a *AzureDevOps) SetPullRequestAutoComplete(projectName string, repositoryName string, pullRequestId int, userId *uuid.UUID) error {
	client, err := git.NewClient(a.ctx, a.connection)
	if err != nil {
		return err
	}

	id := userId.String()
	deleteSourceBranch := true

	_, err = client.UpdatePullRequest(a.ctx, git.UpdatePullRequestArgs{
		Project:       &projectName,
		PullRequestId: &pullRequestId,
		RepositoryId:  &repositoryName,
		GitPullRequestToUpdate: &git.GitPullRequest{
			AutoCompleteSetBy: &webapi.IdentityRef{
				Id: &id,
			},
			CompletionOptions: &git.GitPullRequestCompletionOptions{
				DeleteSourceBranch: &deleteSourceBranch,
				MergeStrategy:      &git.GitPullRequestMergeStrategyValues.Squash,
			},
		},
	})

	return err
}
