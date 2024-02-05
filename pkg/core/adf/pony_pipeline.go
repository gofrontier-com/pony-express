package adf

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyPipeline) AddDependency(pipeline PonyResource) {
	p.Dependencies = append(p.Dependencies, pipeline)
}

func (p *PonyPipeline) GetDependencies() []PonyResource {
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

func FetchPipeline(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]PonyResource, error) {
	result := make([]PonyResource, 0)

	pager := clientFactory.NewPipelinesClient().NewListByFactoryPager(resourceGroup, factoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			p := &PonyPipeline{
				Pipeline: v,
			}
			result = append(result, p)
		}
	}

	return result, nil
}

func GetDependantPipelineNames(p *armdatafactory.PipelineResource) (*[]string, error) {
	var pipelineNames []string
	for _, a := range p.Properties.Activities {
		if *a.GetActivity().Type == "ExecutePipeline" {
			if act, ok := a.(*armdatafactory.ExecutePipelineActivity); ok {
				pipelineNames = append(pipelineNames, *act.TypeProperties.Pipeline.ReferenceName)
			} else {
				return nil, errors.New("Not a ExecutePipelineActivity")
			}
		}
	}
	return &pipelineNames, nil
}

func (a *AzureADFConfig) LoadPipeline(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	ls := &armdatafactory.PipelineResource{}
	ls.UnmarshalJSON(b)
	p := &PonyPipeline{
		Pipeline: ls,
	}
	a.Pipeline = append(a.Pipeline, p)
	return nil
}

func (a *AzureADFConfig) FetchPipeline() error {
	result := make([]PonyResource, 0)
	pipelines, err := FetchPipeline(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	for _, p := range pipelines {
		result = append(result, p)
	}
	a.Pipeline = result
	return nil
}
