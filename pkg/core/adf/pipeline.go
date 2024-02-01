package adf

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func FetchPipeline(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]*PonyPipeline, error) {
	result := make([]*PonyPipeline, 0)

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
	result := make([]*PonyPipeline, 0)
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
