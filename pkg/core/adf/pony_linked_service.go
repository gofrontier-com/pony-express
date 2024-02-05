package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyLinkedService) AddDependency(pipeline PonyResource) {
}

func (p *PonyLinkedService) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyLinkedService) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyLinkedService) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyLinkedService) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyLinkedService) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyLinkedService) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyLinkedService) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyLinkedService) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyLinkedService) GetName() *string {
	return p.LinkedService.Name
}

func (p *PonyLinkedService) ToJSON() []byte {
	bytes, _ := p.LinkedService.MarshalJSON()
	return bytes
}

func (p *PonyLinkedService) FromJSON(bytes []byte) {
	p.LinkedService.UnmarshalJSON(bytes)
}

func FetchLinkedService(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewLinkedServicesClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ls := &PonyLinkedService{
				LinkedService: v,
			}
			result = append(result, ls)
		}
	}

	return result, nil
}

func (a *PonyADF) LoadLinkedService(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	lsr := &armdatafactory.LinkedServiceResource{}
	lsr.UnmarshalJSON(b)
	ls := &PonyLinkedService{
		LinkedService: lsr,
	}
	a.LinkedService = append(a.LinkedService, ls)
	return nil
}

func (a *PonyADF) FetchLinkedService() error {
	ls, err := FetchLinkedService(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.LinkedService = ls
	return nil
}
