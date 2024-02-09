package adf

import (
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (a *PonyADF) AddOrUpdateCredential(p PonyResource) error {
	base := p.Base().(*armdatafactory.ManagedIdentityCredentialResource)
	_, err := a.clientFactory.NewCredentialOperationsClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	if err != nil {
		return err
	}
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateDataset(p PonyResource) error {
	base := p.Base().(*armdatafactory.DatasetResource)
	// _, err := a.clientFactory.NewDatasetsClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateIntegrationRuntime(p PonyResource) error {
	base := p.Base().(*armdatafactory.IntegrationRuntimeResource)
	// _, err := a.clientFactory.NewIntegrationRuntimesClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateLinkedService(p PonyResource) error {
	base := p.Base().(*armdatafactory.LinkedServiceResource)
	// _, err := a.clientFactory.NewLinkedServicesClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateManagedVirtualNetwork(p PonyResource) error {
	base := p.Base().(*armdatafactory.ManagedVirtualNetworkResource)
	// _, err := a.clientFactory.NewManagedVirtualNetworksClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateManagedPrivateEndpoint(p PonyResource) error {
	base := p.Base().(*armdatafactory.ManagedPrivateEndpointResource)
	// _, err := a.clientFactory.NewManagedPrivateEndpointsClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, "default", *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdatePipeline(p PonyResource) error {
	base := p.Base().(*armdatafactory.PipelineResource)
	// _, err := a.clientFactory.NewPipelinesClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateTrigger(p PonyResource) error {
	base := p.Base().(*armdatafactory.TriggerResource)
	// _, err := a.clientFactory.NewTriggersClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), *base, nil)
	// return err
	tb := reflect.TypeOf(base).Elem().Name()
	n := *p.GetName()
	fmt.Printf("Adding or updating %s %s from RG %s and factory %s\n", tb, n, a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) AddOrUpdateFactory() error {
	a.clientFactory.NewFactoriesClient().CreateOrUpdate(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *a.Factory.Base().(*armdatafactory.Factory), nil)
	fmt.Println("updating factory")
	return nil
}
