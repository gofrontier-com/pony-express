package adf

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyCredential) AddDependency(pipeline PonyResource) {
}

func (p *PonyCredential) GetDependencies(resource []PonyResource) []PonyResource {
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

func (a *PonyADF) FetchCredentials() error {
	pager := a.clientFactory.NewCredentialOperationsClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			c := &PonyCredential{
				Credential: v,
			}
			a.Credential = append(a.Credential, c)
		}
	}
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
