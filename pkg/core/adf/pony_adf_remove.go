package adf

import (
	"fmt"
)

func (a *PonyADF) RemoveCredential(p PonyResource) error {
	// _, err := a.clientFactory.NewCredentialOperationsClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing credential %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemoveDataset(p PonyResource) error {
	// _, err := a.clientFactory.NewDatasetsClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing dataset %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemoveIntegrationRuntime(p PonyResource) error {
	// _, err := a.clientFactory.NewIntegrationRuntimesClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing integration runtime %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemoveLinkedService(p PonyResource) error {
	// _, err := a.clientFactory.NewLinkedServicesClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing linked service %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemoveManagedPrivateEndpoint(p PonyResource) error {
	// _, err := a.clientFactory.NewManagedPrivateEndpointsClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing managed private endpoint %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemovePipeline(p PonyResource) error {
	// _, err := a.clientFactory.NewPipelinesClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing pipeline %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}

func (a *PonyADF) RemoveTrigger(p PonyResource) error {
	// _, err := a.clientFactory.NewTriggersClient().Delete(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, *p.GetName(), nil)
	// return err
	fmt.Printf("Removing trigger %s from RG %s and factory %s\n", *p.GetName(), a.Remote.ResourceGroup, a.Remote.FactoryName)
	return nil
}
