package adf

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

type PonyConfig struct {
	Deploy  PonyDeployConfig       `yaml:"deploy"`
	Changes map[string]interface{} `yaml:"changes"`
}

type PonyDeployConfig struct {
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
	Dependencies            []PonyResource
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

type PonyADF struct {
	clientFactory          *armdatafactory.ClientFactory
	ctx                    *context.Context
	Credential             []PonyResource
	Dataset                []PonyResource
	Factory                PonyResource
	IntegrationRuntime     []PonyResource
	LinkedService          []PonyResource
	ManagedVirtualNetwork  []PonyResource
	ManagedPrivateEndpoint []PonyResource
	Pipeline               []PonyResource
	Trigger                []PonyResource
	Remote                 *ADFRemoteConfig
}
