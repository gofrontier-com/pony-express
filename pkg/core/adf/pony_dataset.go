package adf

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyDataset) AddDependency(pipeline PonyResource) {
}

func (p *PonyDataset) GetDependencies(resource []PonyResource) []PonyResource {
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
	pager := a.clientFactory.NewDatasetsClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ds := &PonyDataset{
				Dataset: v,
			}
			a.Dataset = append(a.Dataset, ds)
		}
	}
	return nil
}
