package adf

import "context"

func NewADF(subscriptionId string, resourceGroup string, factoryName string) (*PonyADF, error) {
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
