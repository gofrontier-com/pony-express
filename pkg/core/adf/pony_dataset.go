package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyDataset) AddDependency(pipeline PonyResource) {
}

func (p *PonyDataset) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyDataset) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyDataset) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyDataset) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyDataset) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyDataset) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyDataset) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyDataset) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyDataset) GetName() *string {
	return p.Dataset.Name
}

func (p *PonyDataset) ToJSON() []byte {
	bytes, _ := p.Dataset.MarshalJSON()
	return bytes
}

func (p *PonyDataset) FromJSON(bytes []byte) {
	p.Dataset.UnmarshalJSON(bytes)
}

func FetchDataset(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewDatasetsClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ds := &PonyDataset{
				Dataset: v,
			}
			result = append(result, ds)
		}
	}

	return result, nil
}

func (a *PonyADF) LoadDataset(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	dsr := &armdatafactory.DatasetResource{}
	dsr.UnmarshalJSON(b)
	ds := &PonyDataset{
		Dataset: dsr,
	}
	a.Dataset = append(a.Dataset, ds)
	return nil
}

func (a *PonyADF) FetchDataset() error {
	d, err := FetchDataset(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.Dataset = d
	return nil
}
