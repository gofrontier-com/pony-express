package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyCredential) AddDependency(pipeline PonyResource) {
}

func (p *PonyCredential) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyCredential) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyCredential) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyCredential) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyCredential) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyCredential) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyCredential) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyCredential) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyCredential) GetName() *string {
	return p.Credential.Name
}

func (p *PonyCredential) ToJSON() []byte {
	bytes, _ := p.Credential.MarshalJSON()
	return bytes
}

func (p *PonyCredential) FromJSON(bytes []byte) {
	p.Credential.UnmarshalJSON(bytes)
}

func FetchCredential(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewCredentialOperationsClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			c := &PonyCredential{
				Credential: v,
			}
			result = append(result, c)
		}
	}

	return result, nil
}

func (a *PonyADF) FetchCredentials() error {
	c, err := FetchCredential(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.Credential = c
	return nil
}

func (a *PonyADF) LoadCredential(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	cred := &armdatafactory.ManagedIdentityCredentialResource{}
	cred.UnmarshalJSON(b)
	c := &PonyCredential{
		Credential: cred,
	}
	a.Credential = append(a.Credential, c)
	return nil
}
