package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyManagedVirtualNetwork) AddDependency(pipeline PonyResource) {
}

func (p *PonyManagedVirtualNetwork) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyManagedVirtualNetwork) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyManagedVirtualNetwork) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyManagedVirtualNetwork) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyManagedVirtualNetwork) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyManagedVirtualNetwork) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyManagedVirtualNetwork) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyManagedVirtualNetwork) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyManagedVirtualNetwork) GetName() *string {
	return p.ManagedVirtualNetwork.Name
}

func (p *PonyManagedVirtualNetwork) ToJSON() []byte {
	bytes, _ := p.ManagedVirtualNetwork.MarshalJSON()
	return bytes
}

func (p *PonyManagedVirtualNetwork) FromJSON(bytes []byte) {
	p.ManagedVirtualNetwork.UnmarshalJSON(bytes)
}

func FetchManagedVirtualNetwork(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewManagedVirtualNetworksClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			mvn := &PonyManagedVirtualNetwork{
				ManagedVirtualNetwork: v,
			}
			result = append(result, mvn)
		}
	}
	return result, nil
}

func (a *PonyADF) LoadManagedVirtualNetwork(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	mvnr := &armdatafactory.ManagedVirtualNetworkResource{}
	mvnr.UnmarshalJSON(b)
	mvn := &PonyManagedVirtualNetwork{
		ManagedVirtualNetwork: mvnr,
	}
	a.ManagedVirtualNetwork = append(a.ManagedVirtualNetwork, mvn)
	return nil
}

func (a *PonyADF) FetchManagedVirtualNetwork() error {
	mvn, err := FetchManagedVirtualNetwork(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.ManagedVirtualNetwork = mvn
	return nil
}
