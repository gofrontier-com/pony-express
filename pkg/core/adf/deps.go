package adf

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyPipeline) getPipelineDeps(pipelines []*PonyPipeline) error {
	for _, activity := range p.Pipeline.Properties.Activities {
		if *activity.GetActivity().Type == "ExecutePipeline" {
			act := activity.(*armdatafactory.ExecutePipelineActivity)

			depPipe, err := findMatchingTargetPipeline(act.TypeProperties.Pipeline.ReferenceName, pipelines)
			if err != nil {
				return err
			}

			p.AddDependency(depPipe.Pipeline)
		}
	}
	return nil
}

func (a *AzureADFConfig) Deps() error {
	for _, pipeline := range a.Pipeline {
		if pipeline.RequiresDeployment {
			err := pipeline.getPipelineDeps(a.Pipeline)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
