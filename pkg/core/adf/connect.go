package adf

import (
	"fmt"
)

func Fetch(subscriptionId string, resourceGroup string, factoryName string) (*AzureADFConfig, error) {
	target, err := NewADF(subscriptionId, resourceGroup, factoryName)
	if err != nil {
		return nil, err
	}

	for _, property := range adfFeatures {

		switch property {
		case "factory":
			err := target.FetchFactory()
			if err != nil {
				return nil, err
			}
		case "credential":
			err := target.FetchCredentials()
			if err != nil {
				return nil, err
			}
		case "linkedService":
			err := target.FetchLinkedService()
			if err != nil {
				return nil, err
			}
		case "managedVirtualNetwork":
			err := target.FetchManagedVirtualNetwork()
			if err != nil {
				return nil, err
			}
		case "managedPrivateEndpoint":
			err := target.FetchManagedPrivateEndpoint()
			if err != nil {
				return nil, err
			}
		case "integrationRuntime":
			err := target.FetchIntegrationRuntime()
			if err != nil {
				return nil, err
			}
		case "dataset":
			err := target.FetchDataset()
			if err != nil {
				return nil, err
			}
		case "trigger":
			err := target.FetchTrigger()
			if err != nil {
				return nil, err
			}
		case "pipeline":
			err := target.FetchPipeline()
			if err != nil {
				return nil, err
			}
		default:
			fmt.Println("Not implemented or not used: ", property)
		}
	}

	return target, nil
}
