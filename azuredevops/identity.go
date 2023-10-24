package azuredevops

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/location"
)

// GetIdentityId gets the UUID of the authenticated user. Yes this is weird, see https://github.com/microsoft/azure-devops-python-api/issues/188#issuecomment-494858123
func (a *AzureDevOps) GetIdentityId() (*uuid.UUID, error) {
	client := location.NewClient(a.ctx, a.connection)

	self, err := client.GetConnectionData(a.ctx, location.GetConnectionDataArgs{})
	if err != nil {
		return nil, err
	}

	return self.AuthenticatedUser.Id, nil
}
