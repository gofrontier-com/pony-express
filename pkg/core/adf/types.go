package adf

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// PongConfig is a struct that represents the configuration for the Pony ADF tool
type PonyConfig struct {
	Deploy  PonyDeployConfig       `yaml:"deploy"`
	Changes map[string]interface{} `yaml:"changes"`
}

// PonyDeployConfig is a struct that represents the deployment configuration for the Pony ADF tool.
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

// ADFRemoteConfig is a struct that represents the configuration for the remote ADF
type ADFRemoteConfig struct {
	SubscriptionId string
	ResourceGroup  string
	FactoryName    string
}

// PonyPipeline is a struct that represents an ADF pipeline.
type PonyPipeline struct {
	Pipeline                *armdatafactory.PipelineResource
	Dependencies            []PonyResource
	ConfiguredForDeployment bool
	RequiresDeployment      bool
	ChangeType              int
}

// PonyCredential is a struct that represents an ADF credential.
type PonyCredential struct {
	Credential              *armdatafactory.ManagedIdentityCredentialResource
	ConfiguredForDeployment bool
	RequiresDeployment      bool
	ChangeType              int
}

// PonyFactory is a struct that represents an ADF factory.
type PonyFactory struct {
	Factory                 *armdatafactory.Factory
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyDataset is a struct that represents an ADF dataset.
type PonyDataset struct {
	Dataset                 *armdatafactory.DatasetResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyIntegrationRuntime is a struct that represents an ADF integration runtime.
type PonyIntegrationRuntime struct {
	IntegrationRuntime      *armdatafactory.IntegrationRuntimeResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyLinkedService is a struct that represents an ADF linked service.
type PonyLinkedService struct {
	LinkedService           *armdatafactory.LinkedServiceResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyManagedVirtualNetwork is a struct that represents an ADF managed virtual network.
type PonyManagedVirtualNetwork struct {
	ManagedVirtualNetwork   *armdatafactory.ManagedVirtualNetworkResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyManagedPrivateEndpoint is a struct that represents an ADF managed private endpoint.
type PonyManagedPrivateEndpoint struct {
	ManagedPrivateEndpoint  *armdatafactory.ManagedPrivateEndpointResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyTrigger is a struct that represents an ADF trigger.
type PonyTrigger struct {
	Trigger                 *armdatafactory.TriggerResource
	RequiresDeployment      bool
	ConfiguredForDeployment bool
	ChangeType              int
}

// PonyADF is a struct that represents an ADF.
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
