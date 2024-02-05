package adf

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyFactory) AddDependency(pipeline PonyResource) {
}

func (p *PonyFactory) GetDependencies() []PonyResource {
	return nil
}

func (p *PonyFactory) getPipelineDeps([]PonyResource) error {
	return nil
}

func (p *PonyFactory) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyFactory) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyFactory) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyFactory) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyFactory) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyFactory) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyFactory) GetName() *string {
	return p.Factory.Name
}

func (p *PonyFactory) ToJSON() []byte {
	bytes, _ := p.Factory.MarshalJSON()
	return bytes
}

func (p *PonyFactory) FromJSON(bytes []byte) {
	p.Factory.UnmarshalJSON(bytes)
}

func FetchFactory(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, subscriptionId string, factoryName string) (PonyResource, error) {
	res, err := clientFactory.NewFactoriesClient().Get(*ctx, subscriptionId, factoryName, &armdatafactory.FactoriesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		return nil, err
	}

	return &PonyFactory{Factory: &res.Factory}, err
}

func (a *AzureADFConfig) FetchFactory() error {
	f, err := FetchFactory(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.Factory = f
	return nil
}

func (a *AzureADFConfig) LoadFactory(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	af := &armdatafactory.Factory{}
	af.UnmarshalJSON(b)
	f := &PonyFactory{
		Factory: af,
	}
	a.Factory = f
	return nil
}
