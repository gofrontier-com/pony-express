package adf

import (
	"context"
	"fmt"
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
func (a *PonyADF) Deps() bool {
	depsSatisfied := true

	for _, pipeline := range a.Pipeline {
		if pipeline.GetConfiguredForDeployment() {
			pipeline.GetDependencies(a.Pipeline)
			depsSatisfied = pipeline.CheckDependencies() && depsSatisfied
		}
	}

	for _, ls := range a.LinkedService {
		if ls.GetConfiguredForDeployment() {
			ls.GetDependencies(a.Credential, a.IntegrationRuntime)
			depsSatisfied = ls.CheckDependencies() && depsSatisfied
		}
	}

	for _, ds := range a.Dataset {
		if ds.GetConfiguredForDeployment() {
			ds.GetDependencies(a.LinkedService)
			depsSatisfied = ds.CheckDependencies() && depsSatisfied
		}
	}

	for _, t := range a.Trigger {
		if t.GetConfiguredForDeployment() {
			t.GetDependencies(a.Pipeline)
			depsSatisfied = t.CheckDependencies() && depsSatisfied
		}
	}

	return depsSatisfied
}

// Diff compares two PonyADF objects.
func (a *PonyADF) Diff(target *PonyADF) {
	compareFactory(a.Factory, target.Factory, "Factory.Identity", "Factory.Properties.PublicNetworkAccess", "Factory.Properties.ProvisioningState", "Factory.Properties.CreateTime", "Factory.Properties.Version")
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
