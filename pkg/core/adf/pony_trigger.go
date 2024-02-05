package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyTrigger) AddDependency(pipeline PonyResource) {
}

func (p *PonyTrigger) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyTrigger) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyTrigger) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyTrigger) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyTrigger) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyTrigger) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyTrigger) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyTrigger) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyTrigger) GetName() *string {
	return p.Trigger.Name
}

func (p *PonyTrigger) ToJSON() []byte {
	bytes, _ := p.Trigger.MarshalJSON()
	return bytes
}

func (p *PonyTrigger) FromJSON(bytes []byte) {
	p.Trigger.UnmarshalJSON(bytes)
}

func FetchTrigger(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewTriggersClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			t := &PonyTrigger{
				Trigger: v,
			}
			result = append(result, t)
		}
	}

	return result, nil
}

func (a *AzureADFConfig) FetchTrigger() error {
	t, err := FetchTrigger(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.Trigger = t
	return nil
}

func (a *AzureADFConfig) LoadTrigger(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	trigger := &armdatafactory.TriggerResource{}
	trigger.UnmarshalJSON(b)
	t := &PonyTrigger{
		Trigger: trigger,
	}
	a.Trigger = append(a.Trigger, t)
	return nil
}
