package adf

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Fetch fetches the ADF resources from the remote ADF.
func (a *PonyADF) Fetch() error {
	err := a.FetchFactory()
	if err != nil {
		return err
	}

	err = a.FetchCredentials()
	if err != nil {
		return err
	}

	err = a.FetchLinkedService()
	if err != nil {
		return err
	}

	err = a.FetchManagedVirtualNetwork()
	if err != nil {
		return err
	}

	err = a.FetchManagedPrivateEndpoint()
	if err != nil {
		return err
	}

	err = a.FetchIntegrationRuntime()
	if err != nil {
		return err
	}

	err = a.FetchDataset()
	if err != nil {
		return err
	}

	err = a.FetchTrigger()
	if err != nil {
		return err
	}

	err = a.FetchPipeline()
	if err != nil {
		return err
	}

	return nil
}

// FetchCredentials fetches the credentials from the remote ADF.
func (a *PonyADF) FetchCredentials() error {
	pager := a.clientFactory.NewCredentialOperationsClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			c := &PonyCredential{
				Credential: v,
			}
			a.Credential = append(a.Credential, c)
		}
	}
	return nil
}

// FetchDataset fetches the datasets from the remote ADF.
func (a *PonyADF) FetchDataset() error {
	pager := a.clientFactory.NewDatasetsClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ds := &PonyDataset{
				Dataset: v,
			}
			a.Dataset = append(a.Dataset, ds)
		}
	}
	return nil
}

// FetchFactory fetches the factory from the remote ADF.
func (a *PonyADF) FetchFactory() error {
	res, err := a.clientFactory.NewFactoriesClient().Get(*a.ctx, a.Remote.ResourceGroup, a.Remote.FactoryName, &armdatafactory.FactoriesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		return err
	}
	a.Factory = &PonyFactory{Factory: &res.Factory}
	return nil
}

// FetchIntegrationRuntime fetches the integration runtimes from the remote ADF.
func (a *PonyADF) FetchIntegrationRuntime() error {
	pager := a.clientFactory.NewIntegrationRuntimesClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ir := &PonyIntegrationRuntime{
				IntegrationRuntime: v,
			}
			a.IntegrationRuntime = append(a.IntegrationRuntime, ir)
		}
	}
	return nil
}

// FetchLinkedService fetches the linked services from the remote ADF.
func (a *PonyADF) FetchLinkedService() error {
	pager := a.clientFactory.NewLinkedServicesClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			ls := &PonyLinkedService{
				LinkedService: v,
			}
			a.LinkedService = append(a.LinkedService, ls)
		}
	}
	return nil
}

// FetchManagedPrivateEndpoint fetches the managed private endpoints from the remote ADF.
func (a *PonyADF) FetchManagedPrivateEndpoint() error {
	pager := a.clientFactory.NewManagedPrivateEndpointsClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, "default", nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			mpe := &PonyManagedPrivateEndpoint{
				ManagedPrivateEndpoint: v,
			}
			a.ManagedPrivateEndpoint = append(a.ManagedPrivateEndpoint, mpe)
		}
	}
	return nil
}

// FetchManagedVirtualNetwork fetches the managed virtual networks from the remote ADF.
func (a *PonyADF) FetchManagedVirtualNetwork() error {
	pager := a.clientFactory.NewManagedVirtualNetworksClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			mvn := &PonyManagedVirtualNetwork{
				ManagedVirtualNetwork: v,
			}
			a.ManagedVirtualNetwork = append(a.ManagedVirtualNetwork, mvn)
		}
	}
	return nil
}

// FetchPipeline fetches the pipelines from the remote ADF.
func (a *PonyADF) FetchPipeline() error {
	pager := a.clientFactory.NewPipelinesClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			p := &PonyPipeline{
				Pipeline: v,
			}
			a.Pipeline = append(a.Pipeline, p)
		}
	}
	return nil
}

// FetchTrigger fetches the triggers from the remote ADF.
func (a *PonyADF) FetchTrigger() error {
	pager := a.clientFactory.NewTriggersClient().NewListByFactoryPager(a.Remote.ResourceGroup, a.Remote.FactoryName, nil)

	for pager.More() {
		page, err := pager.NextPage(*a.ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}

		for _, v := range page.Value {
			t := &PonyTrigger{
				Trigger: v,
			}
			a.Trigger = append(a.Trigger, t)
		}
	}
	return nil
}
