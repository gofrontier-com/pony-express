package adf

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyIntegrationRuntime) AddDependency(pipeline PonyResource) {
}

func (p *PonyIntegrationRuntime) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyIntegrationRuntime) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyIntegrationRuntime) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyIntegrationRuntime) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyIntegrationRuntime) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyIntegrationRuntime) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyIntegrationRuntime) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyIntegrationRuntime) GetName() *string {
	return p.IntegrationRuntime.Name
}

func (p *PonyIntegrationRuntime) ToJSON() []byte {
	bytes, _ := p.IntegrationRuntime.MarshalJSON()
	return bytes
}

func (p *PonyIntegrationRuntime) FromJSON(bytes []byte) {
	p.IntegrationRuntime.UnmarshalJSON(bytes)
}

func (a *PonyADF) LoadIntegrationRuntime(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	irr := &armdatafactory.IntegrationRuntimeResource{}
	irr.UnmarshalJSON(b)
	ir := &PonyIntegrationRuntime{
		IntegrationRuntime: irr,
	}
	a.IntegrationRuntime = append(a.IntegrationRuntime, ir)
	return nil
}

func (a *PonyADF) FetchIntegrationRuntime() error {
	pager := a.clientFactory.NewIntegrationRuntimesClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ir := &PonyIntegrationRuntime{
				IntegrationRuntime: v,
			}
			a.IntegrationRuntime = append(a.IntegrationRuntime, ir)
		}
	}
	return nil
}
