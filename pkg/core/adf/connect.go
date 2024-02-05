package adf

func Fetch(subscriptionId string, resourceGroup string, factoryName string) (*PonyADF, error) {
	target, err := NewADF(subscriptionId, resourceGroup, factoryName)
	if err != nil {
		return nil, err
	}

	err = target.FetchFactory()
	if err != nil {
		return nil, err
	}

	err = target.FetchCredentials()
	if err != nil {
		return nil, err
	}

	err = target.FetchLinkedService()
	if err != nil {
		return nil, err
	}

	err = target.FetchManagedVirtualNetwork()
	if err != nil {
		return nil, err
	}

	err = target.FetchManagedPrivateEndpoint()
	if err != nil {
		return nil, err
	}

	err = target.FetchIntegrationRuntime()
	if err != nil {
		return nil, err
	}

	err = target.FetchDataset()
	if err != nil {
		return nil, err
	}

	err = target.FetchTrigger()
	if err != nil {
		return nil, err
	}

	err = target.FetchPipeline()
	if err != nil {
		return nil, err
	}

	return target, nil
}
