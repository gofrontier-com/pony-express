package adf

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func FetchFactory(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, subscriptionId string, factoryName string) (*PonyFactory, error) {
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
