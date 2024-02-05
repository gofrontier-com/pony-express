package adf

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

func NewADF() *PonyADF {
	return &PonyADF{}
}

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

func (a *PonyADF) ProcessChanges(adfChanges map[string]interface{}) error {
	changes, err := getJsonPatches(adfChanges)
	if err != nil {
		return err
	}

	a.processChanges(changes)
	return nil
}

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

func (a *PonyADF) Deps() error {
	for _, pipeline := range a.Pipeline {
		if pipeline.GetRequiresDeployment() {
			pipeline.GetDependencies(a.Pipeline)
		}
	}
	return nil
}

func (a *PonyADF) Diff(target *PonyADF) {
	compareFactory(a.Factory, target.Factory)
	compare(a.Credential, target.Credential)
	compare(a.LinkedService, target.LinkedService)
	compare(a.ManagedVirtualNetwork, target.ManagedVirtualNetwork)
	compare(a.ManagedPrivateEndpoint, target.ManagedPrivateEndpoint)
	compare(a.IntegrationRuntime, target.IntegrationRuntime)
	compare(a.Dataset, target.Dataset)
	compare(a.Trigger, target.Trigger)
	compare(a.Pipeline, target.Pipeline)
}
