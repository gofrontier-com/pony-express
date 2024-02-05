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
