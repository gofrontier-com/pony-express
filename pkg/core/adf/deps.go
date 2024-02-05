package adf

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyPipeline) getPipelineDeps(pipelines []PonyResource) error {
	for _, activity := range p.Pipeline.Properties.Activities {
		if *activity.GetActivity().Type == "ExecutePipeline" {
			act := activity.(*armdatafactory.ExecutePipelineActivity)

			depPipe, err := findMatchingTarget(act.TypeProperties.Pipeline.ReferenceName, pipelines)
			if err != nil {
				return err
			}

			p.AddDependency(depPipe)
		}
	}
	return nil
}

func (a *PonyADF) Deps() error {
	for _, pipeline := range a.Pipeline {
		if pipeline.GetRequiresDeployment() {
			err := pipeline.getPipelineDeps(a.Pipeline)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
