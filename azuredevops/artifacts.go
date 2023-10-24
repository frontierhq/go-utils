package azuredevops

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/feed"
)

// GetPackageVersion gets all GitRepository
func (a *AzureDevOps) GetPackageVersion(projectName string, feedName string) (*[]feed.Package, error) {
	feedClient, err := feed.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, err
	}

	getPackagesArgs := feed.GetPackagesArgs{
		FeedId:  &feedName,
		Project: &projectName,
	}
	return feedClient.GetPackages(a.ctx, getPackagesArgs)
}
