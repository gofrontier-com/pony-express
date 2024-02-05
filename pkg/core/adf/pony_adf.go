package adf

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// NewRemoteADF creates a new PonyADF object for a remote ADF.
func NewRemoteADF(subscriptionId string, resourceGroup string, factoryName string) (*PonyADF, error) {
	clientFactory, err := CreateClientFactory(subscriptionId)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &PonyADF{
		clientFactory: clientFactory,
		ctx:           &ctx,
		Remote: &ADFRemoteConfig{
			SubscriptionId: subscriptionId,
			ResourceGroup:  resourceGroup,
			FactoryName:    factoryName,
		},
	}, err
}

// NewADF creates a new PonyADF object.
func NewADF() *PonyADF {
	return &PonyADF{}
}

// LoadFromFolder loads the ADF resources from the specified folder.
func (a *PonyADF) LoadFromFolder(configDir string) error {
	err := filepath.Walk(configDir, func(path string, f os.FileInfo, err error) error {
		s := strings.Replace(path, configDir+string(filepath.Separator), "", 1)
		ss := strings.Split(s, string(filepath.Separator))

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			if len(ss) > 1 {
				property := ss[len(ss)-2]
				switch property {
				case "factory":
					err = a.LoadFactory(path)
				case "credential":
					err = a.LoadCredential(path)
				case "linkedService":
					err = a.LoadLinkedService(path)
				case "managedVirtualNetwork":
					err = a.LoadManagedVirtualNetwork(path)
				case "managedPrivateEndpoint":
					err = a.LoadManagedPrivateEndPoint(path)
				case "integrationRuntime":
					err = a.LoadIntegrationRuntime(path)
				case "dataset":
					err = a.LoadDataset(path)
				case "trigger":
					err = a.LoadTrigger(path)
				case "pipeline":
					err = a.LoadPipeline(path)
				default:
					fmt.Println("Not implemented or not used: ", property)
				}
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	return err
}

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

func (a *PonyADF) processChanges(changes []*Change) {
	for _, c := range changes {
		switch c.Type {
		case "pipeline":
			processChanges(c, a.Pipeline)
		case "dataset":
			processChanges(c, a.Dataset)
		case "linkedService":
			processChanges(c, a.LinkedService)
		case "integrationRuntime":
			processChanges(c, a.IntegrationRuntime)
		case "managedVirtualNetwork":
			processChanges(c, a.ManagedVirtualNetwork)
		case "managedPrivateEndpoint":
			processChanges(c, a.ManagedPrivateEndpoint)
		case "factory":
			processChange(c, a.Factory)
		case "trigger":
			processChanges(c, a.Trigger)
		case "credential":
			processChanges(c, a.Credential)
		default:
			fmt.Println("Unknown change type: ", c.Type)
		}
	}
}

// ProcessChanges processes the changes to the ADF resources.
func (a *PonyADF) ProcessChanges(adfChanges map[string]interface{}) error {
	changes, err := getJsonPatches(adfChanges)
	if err != nil {
		return err
	}

	a.processChanges(changes)
	return nil
}

// SetDeploymentConfig sets the deployment configuration for the ADF resources.
func (a *PonyADF) SetDeploymentConfig(config *PonyDeployConfig) {
	setDeploymentConfig(config.Credential, a.Credential)
	setDeploymentConfig(config.Pipeline, a.Pipeline)
	setDeploymentConfig(config.Trigger, a.Trigger)
	setDeploymentConfig(config.Dataset, a.Dataset)
	setDeploymentConfig(config.IntegrationRuntime, a.IntegrationRuntime)
	setDeploymentConfig(config.LinkedService, a.LinkedService)
	setDeploymentConfig(config.ManagedPrivateEndpoint, a.ManagedPrivateEndpoint)
	setDeploymentConfig(config.ManagedVirtualNetwork, a.ManagedVirtualNetwork)
}

// SetTargetDeploymentConfig sets the target deployment configuration for the ADF resources.
func (a *PonyADF) SetTargetDeploymentConfig(config *PonyDeployConfig) {
	setTargetDeploymentConfig(config.Credential, a.Credential)
	setTargetDeploymentConfig(config.Pipeline, a.Pipeline)
	setTargetDeploymentConfig(config.Trigger, a.Trigger)
	setTargetDeploymentConfig(config.Dataset, a.Dataset)
	setTargetDeploymentConfig(config.IntegrationRuntime, a.IntegrationRuntime)
	setTargetDeploymentConfig(config.LinkedService, a.LinkedService)
	setTargetDeploymentConfig(config.ManagedPrivateEndpoint, a.ManagedPrivateEndpoint)
	setTargetDeploymentConfig(config.ManagedVirtualNetwork, a.ManagedVirtualNetwork)
}

// Deps sets the dependencies for the ADF resources.
func (a *PonyADF) Deps() error {
	for _, pipeline := range a.Pipeline {
		if pipeline.GetRequiresDeployment() {
			pipeline.GetDependencies(a.Pipeline)
		}
	}
	return nil
}

// Diff compares two PonyADF objects.
func (a *PonyADF) Diff(target *PonyADF) {
	compareFactory(a.Factory, target.Factory)
	compare(a.Credential, target.Credential, "Credential")
	compare(a.LinkedService, target.LinkedService, "LinkedService")
	compare(a.ManagedVirtualNetwork, target.ManagedVirtualNetwork, "ManagedVirtualNetwork",
		"ManagedVirtualNetwork.Properties")
	compare(a.ManagedPrivateEndpoint, target.ManagedPrivateEndpoint, "ManagedPrivateEndpoint",
		"ManagedPrivateEndpoint.Properties.AdditionalProperties",
		"ManagedPrivateEndpoint.Properties.ConnectionState",
		"ManagedPrivateEndpoint.Properties.Fqdns",
		"ManagedPrivateEndpoint.Properties.ProvisioningState")
	compare(a.IntegrationRuntime, target.IntegrationRuntime, "IntegrationRuntime")
	compare(a.Dataset, target.Dataset, "Dataset")
	compare(a.Trigger, target.Trigger, "Trigger")
	compare(a.Pipeline, target.Pipeline, "Pipeline")
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

// LoadCredential loads the credential from the specified file.
func (a *PonyADF) LoadCredential(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	cred := &armdatafactory.ManagedIdentityCredentialResource{}
	cred.UnmarshalJSON(b)
	c := &PonyCredential{
		Credential: cred,
	}
	a.Credential = append(a.Credential, c)
	return nil
}

// LoadDataset loads the dataset from the specified file.
func (a *PonyADF) LoadDataset(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	dsr := &armdatafactory.DatasetResource{}
	dsr.UnmarshalJSON(b)
	ds := &PonyDataset{
		Dataset: dsr,
	}
	a.Dataset = append(a.Dataset, ds)
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

// LoadFactory loads the factory from the specified file.
func (a *PonyADF) LoadFactory(filePath string) error {
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

// LoadIntegrationRuntime loads the integration runtime from the specified file.
func (a *PonyADF) LoadIntegrationRuntime(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	irr := &armdatafactory.IntegrationRuntimeResource{}
	irr.UnmarshalJSON(b)
	ir := &PonyIntegrationRuntime{
		IntegrationRuntime: irr,
	}
	a.IntegrationRuntime = append(a.IntegrationRuntime, ir)
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

// LoadLinkedService loads the linked service from the specified file.
func (a *PonyADF) LoadLinkedService(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	lsr := &armdatafactory.LinkedServiceResource{}
	lsr.UnmarshalJSON(b)
	ls := &PonyLinkedService{
		LinkedService: lsr,
	}
	a.LinkedService = append(a.LinkedService, ls)
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

// LoadManagedPrivateEndPoint loads the managed private endpoint from the specified file.
func (a *PonyADF) LoadManagedPrivateEndPoint(filePath string) error {
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

// LoadManagedVirtualNetwork loads the managed virtual network from the specified file.
func (a *PonyADF) LoadManagedVirtualNetwork(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	mvnr := &armdatafactory.ManagedVirtualNetworkResource{}
	mvnr.UnmarshalJSON(b)
	mvn := &PonyManagedVirtualNetwork{
		ManagedVirtualNetwork: mvnr,
	}
	a.ManagedVirtualNetwork = append(a.ManagedVirtualNetwork, mvn)
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

// LoadPipeline loads the pipeline from the specified file.
func (a *PonyADF) LoadPipeline(filePath string) error {
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

// LoadTrigger loads the trigger from the specified file.
func (a *PonyADF) LoadTrigger(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	trigger := &armdatafactory.TriggerResource{}
	trigger.UnmarshalJSON(b)
	t := &PonyTrigger{
		Trigger: trigger,
	}
	a.Trigger = append(a.Trigger, t)
	return nil
}
