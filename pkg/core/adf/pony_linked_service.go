package adf

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyLinkedService) AddDependency(pipeline PonyResource) {
}

func (p *PonyLinkedService) GetDependencies(resource []PonyResource) []PonyResource {
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
	pager := a.clientFactory.NewLinkedServicesClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ls := &PonyLinkedService{
				LinkedService: v,
			}
			a.LinkedService = append(a.LinkedService, ls)
		}
	}
	return nil
}
