package adf

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

const (
	Add = iota
	Update
	Remove
)

type AppADFConfig struct {
	Deploy  ADFDeployConfig        `yaml:"deploy"`
	Changes map[string]interface{} `yaml:"changes"`
}

type ADFDeployConfig struct {
	Credential             []string `yaml:"credential"`
	Dataset                []string `yaml:"dataset"`
	Factory                []string `yaml:"factory"`
	IntegrationRuntime     []string `yaml:"integrationRuntime"`
	LinkedService          []string `yaml:"linkedService"`
	ManagedVirtualNetwork  []string `yaml:"managedVirtualNetwork"`
	ManagedPrivateEndpoint []string `yaml:"managedPrivateEndpoint"`
	Pipeline               []string `yaml:"pipeline"`
	Trigger                []string `yaml:"trigger"`
}

type ADFRemoteConfig struct {
	SubscriptionId string
	ResourceGroup  string
	FactoryName    string
}

type PonyPipeline struct {
	Pipeline                *armdatafactory.PipelineResource
	Dependencies            []*armdatafactory.PipelineResource
	ConfiguredForDeployment bool
	RequiresDeployment      bool
	ChangeType              int
}

type PonyCredential struct {
	Credential              *armdatafactory.ManagedIdentityCredentialResource
	ConfiguredForDeployment bool
	RequiresDeployment      bool
	ChangeType              int
}

func (p *PonyPipeline) AddDependency(pipeline *armdatafactory.PipelineResource) {
	p.Dependencies = append(p.Dependencies, pipeline)
}

type PonyFactory struct {
	Factory                 *armdatafactory.Factory
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyDataset struct {
	Dataset                 *armdatafactory.DatasetResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyIntegrationRuntime struct {
	IntegrationRuntime      *armdatafactory.IntegrationRuntimeResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyLinkedService struct {
	LinkedService           *armdatafactory.LinkedServiceResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyManagedVirtualNetwork struct {
	ManagedVirtualNetwork   *armdatafactory.ManagedVirtualNetworkResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyManagedPrivateEndpoint struct {
	ManagedPrivateEndpoint  *armdatafactory.ManagedPrivateEndpointResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type PonyTrigger struct {
	Trigger                 *armdatafactory.TriggerResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

type AzureADFConfig struct {
	clientFactory          *armdatafactory.ClientFactory
	ctx                    *context.Context
	Credential             []*PonyCredential
	Dataset                []*PonyDataset
	Factory                *PonyFactory
	IntegrationRuntime     []*PonyIntegrationRuntime
	LinkedService          []*PonyLinkedService
	ManagedVirtualNetwork  []*PonyManagedVirtualNetwork
	ManagedPrivateEndpoint []*PonyManagedPrivateEndpoint
	Pipeline               []*PonyPipeline
	Trigger                []*PonyTrigger
	Remote                 *ADFRemoteConfig
}

func NewADF(subscriptionId string, resourceGroup string, factoryName string) (*AzureADFConfig, error) {
	clientFactory, err := CreateClientFactory(subscriptionId)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &AzureADFConfig{
		clientFactory: clientFactory,
		ctx:           &ctx,
		Remote: &ADFRemoteConfig{
			SubscriptionId: subscriptionId,
			ResourceGroup:  resourceGroup,
			FactoryName:    factoryName,
		},
	}, err
}
