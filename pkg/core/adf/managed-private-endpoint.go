package adf

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func FetchManagedPrivateEndpoint(clientFactory *armdatafactory.ClientFactory, ctx *context.Context, resourceGroup string, factoryName string) ([]*PonyManagedPrivateEndpoint, error) {
	result := make([]*PonyManagedPrivateEndpoint, 0)

	pager := clientFactory.NewManagedPrivateEndpointsClient().NewListByFactoryPager(resourceGroup, factoryName, "default", nil)

	for pager.More() {
		page, err := pager.NextPage(*ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			mpe := &PonyManagedPrivateEndpoint{
				ManagedPrivateEndpoint: v,
			}
			result = append(result, mpe)
		}
	}

	return result, nil
}

func (a *AzureADFConfig) LoadManagedPrivateEndPoint(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	mper := &armdatafactory.ManagedPrivateEndpointResource{}
	mper.UnmarshalJSON(b)
	mpe := &PonyManagedPrivateEndpoint{
		ManagedPrivateEndpoint: mper,
	}
	a.ManagedPrivateEndpoint = append(a.ManagedPrivateEndpoint, mpe)
	return nil
}

func (a *AzureADFConfig) FetchManagedPrivateEndpoint() error {
	mpe, err := FetchManagedPrivateEndpoint(a.clientFactory, a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName)
	if err != nil {
		return err
	}
	a.ManagedPrivateEndpoint = mpe
	return nil
}
