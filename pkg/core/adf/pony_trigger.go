package adf

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyTrigger) AddDependency(pipeline PonyResource) {
	p.Dependencies = append(p.Dependencies, pipeline)
}

func (p *PonyTrigger) Base() interface{} {
	return p.Trigger
}

func (p *PonyTrigger) CheckDependencies() bool {
	if !p.ConfiguredForDeployment {
		return true
	}

	for _, dep := range p.Dependencies {
		if !dep.GetConfiguredForDeployment() {
			fmt.Println("Trigger ", *p.GetName(), " has a dependency that is not configured for deployment: ", *dep.GetName())
			return false
		}
	}
	return true
}

func (p *PonyTrigger) GetDependencies(resource ...[]PonyResource) []PonyResource {
	depPipes := p.Trigger.Properties.(*armdatafactory.ScheduleTrigger).Pipelines
	for _, depPipe := range depPipes {
		for _, pipeline := range resource[0] {
			if *pipeline.GetName() == *depPipe.PipelineReference.ReferenceName {
				p.AddDependency(pipeline)
			}
		}
	}
	return p.Dependencies
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
