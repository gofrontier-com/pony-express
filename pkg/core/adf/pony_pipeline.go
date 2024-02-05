package adf

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyPipeline) AddDependency(pipeline PonyResource) {
	p.Dependencies = append(p.Dependencies, pipeline)
}

func (p *PonyPipeline) GetDependencies(pipelines []PonyResource) []PonyResource {
	if len(p.Dependencies) > 0 {
		return p.Dependencies
	}

	for _, activity := range p.Pipeline.Properties.Activities {
		if *activity.GetActivity().Type == "ExecutePipeline" {
			act := activity.(*armdatafactory.ExecutePipelineActivity)

			depPipe, err := findMatchingTarget(act.TypeProperties.Pipeline.ReferenceName, pipelines)
			if err != nil {
				continue
			}

			p.AddDependency(depPipe)
		}
	}
	return p.Dependencies
}

func (p *PonyPipeline) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyPipeline) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyPipeline) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyPipeline) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyPipeline) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyPipeline) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyPipeline) GetName() *string {
	return p.Pipeline.Name
}

func (p *PonyPipeline) ToJSON() []byte {
	bytes, _ := p.Pipeline.MarshalJSON()
	return bytes
}

func (p *PonyPipeline) FromJSON(bytes []byte) {
	p.Pipeline.UnmarshalJSON(bytes)
}
