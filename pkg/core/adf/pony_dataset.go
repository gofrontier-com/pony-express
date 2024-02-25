package adf

import (
	"fmt"
	"reflect"
)

func (p *PonyDataset) getDependantLinkedServices(linkedServices []PonyResource) {
	lsName := p.Dataset.Properties.GetDataset().LinkedServiceName.ReferenceName

	for _, ls := range linkedServices {
		if *ls.GetName() == *lsName {
			p.AddDependency(ls)
		}
	}
}

func (p *PonyDataset) AddDependency(resource PonyResource) {
	p.Dependencies = append(p.Dependencies, resource)
}

func (p *PonyDataset) Base() interface{} {
	return p.Dataset
}

func (p *PonyDataset) CheckDependencies() bool {
	if !p.ConfiguredForDeployment {
		return true
	}

	depsSatisfied := true
	for _, dep := range p.Dependencies {
		if !dep.GetConfiguredForDeployment() {
			fmt.Println("Dataset ", *p.GetName(), " has a dependency that is not configured for deployment: ", *dep.GetName())
			depsSatisfied = false
		}
	}

	return depsSatisfied
}

func (p *PonyDataset) GetDependencies(resources ...[]PonyResource) []PonyResource {
	var linkedServices []PonyResource

	for i, resource := range resources {
		for _, r := range resource {
			t := reflect.TypeOf(r).Elem().Name()
			if t == "PonyLinkedService" {
				linkedServices = resources[i]
			}
			break
		}
	}

	p.getDependantLinkedServices(linkedServices)

	return p.Dependencies
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
