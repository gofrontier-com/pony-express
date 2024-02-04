package adf

import (
	"fmt"
	"strings"

	"github.com/go-test/deep"
)

func filterRemoteProps(prop string, extraProps ...string) bool {
	props := []string{"Etag", "ID", "Type"}
	if extraProps != nil {
		props = append(props, extraProps...)
	}
	for _, p := range props {
		if prop == p {
			return false
		}
	}
	return true
}

func compareFactory(source *PonyFactory, target *PonyFactory) {
	if diff := deep.Equal(source.Factory, target.Factory); diff != nil {
		for _, d := range diff {
			prop := strings.Split(d, ":")[0]
			if filterRemoteProps(prop) {
				source.RequiresDeployment = true
			}
		}
	}

}

func findMatchingTargetCredential(source *PonyCredential, target []*PonyCredential) (*PonyCredential, error) {
	for _, t := range target {
		if *source.Credential.Name == *t.Credential.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.Credential.Name)
}

func compareCredentials(source []*PonyCredential, target []*PonyCredential) {
	for _, s := range source {
		t, err := findMatchingTargetCredential(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.Credential, t.Credential); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetCredential(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetIntegrationRuntime(source *PonyIntegrationRuntime, target []*PonyIntegrationRuntime) (*PonyIntegrationRuntime, error) {
	for _, t := range target {
		if *source.IntegrationRuntime.Name == *t.IntegrationRuntime.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.IntegrationRuntime.Name)
}

func compareIntegrationRuntimes(source []*PonyIntegrationRuntime, target []*PonyIntegrationRuntime) {
	for _, s := range source {
		t, err := findMatchingTargetIntegrationRuntime(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.IntegrationRuntime, t.IntegrationRuntime); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetIntegrationRuntime(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetLinkedService(source *PonyLinkedService, target []*PonyLinkedService) (*PonyLinkedService, error) {
	for _, t := range target {
		if *source.LinkedService.Name == *t.LinkedService.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.LinkedService.Name)
}

func compareLinkedServices(source []*PonyLinkedService, target []*PonyLinkedService) {
	for _, s := range source {
		t, err := findMatchingTargetLinkedService(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.LinkedService, t.LinkedService); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetLinkedService(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetManagedVirtualNetwork(source *PonyManagedVirtualNetwork, target []*PonyManagedVirtualNetwork) (*PonyManagedVirtualNetwork, error) {
	for _, t := range target {
		if *source.ManagedVirtualNetwork.Name == *t.ManagedVirtualNetwork.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.ManagedVirtualNetwork.Name)
}

func compareManagedVirtualNetworks(source []*PonyManagedVirtualNetwork, target []*PonyManagedVirtualNetwork) {
	for _, s := range source {
		t, err := findMatchingTargetManagedVirtualNetwork(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.ManagedVirtualNetwork, t.ManagedVirtualNetwork); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop, "Properties") {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetManagedVirtualNetwork(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetManagedPrivateEndpoint(source *PonyManagedPrivateEndpoint, target []*PonyManagedPrivateEndpoint) (*PonyManagedPrivateEndpoint, error) {
	for _, t := range target {
		if *source.ManagedPrivateEndpoint.Name == *t.ManagedPrivateEndpoint.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.ManagedPrivateEndpoint.Name)
}

func compareManagedPrivateEndpoints(source []*PonyManagedPrivateEndpoint, target []*PonyManagedPrivateEndpoint) {
	for _, s := range source {
		t, err := findMatchingTargetManagedPrivateEndpoint(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.ManagedPrivateEndpoint, t.ManagedPrivateEndpoint); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop, "Properties.AdditionalProperties", "Properties.ConnectionState", "Properties.Fqdns", "Properties.ProvisioningState") {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetManagedPrivateEndpoint(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetDataset(source *PonyDataset, target []*PonyDataset) (*PonyDataset, error) {
	for _, t := range target {
		if *source.Dataset.Name == *t.Dataset.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.Dataset.Name)
}

func compareDatasets(source []*PonyDataset, target []*PonyDataset) {
	for _, s := range source {
		t, err := findMatchingTargetDataset(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.Dataset, t.Dataset); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetDataset(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetTrigger(source *PonyTrigger, target []*PonyTrigger) (*PonyTrigger, error) {
	for _, t := range target {
		if *source.Trigger.Name == *t.Trigger.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *source.Trigger.Name)
}

func compareTriggers(source []*PonyTrigger, target []*PonyTrigger) {
	for _, s := range source {
		t, err := findMatchingTargetTrigger(s, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.Trigger, t.Trigger); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetTrigger(t, source)
		if err != nil {
			t.RequiresDeployment = true
			t.ChangeType = Remove
		}
	}
}

func findMatchingTargetPipeline(sourceName *string, target []*PonyPipeline) (*PonyPipeline, error) {
	for _, t := range target {
		if *sourceName == *t.Pipeline.Name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *sourceName)
}

func comparePipelines(source []*PonyPipeline, target []*PonyPipeline) {
	for _, s := range source {
		s.RequiresDeployment = false
		t, err := findMatchingTargetPipeline(s.Pipeline.Name, target)
		if err != nil {
			s.RequiresDeployment = true
			s.ChangeType = Add
			continue
		}

		if diff := deep.Equal(s.Pipeline, t.Pipeline); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.RequiresDeployment = true
					s.ChangeType = Update
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTargetPipeline(t.Pipeline.Name, source)
		if err != nil {
			t.RequiresDeployment = true
			// t.ConfiguredForDeployment = true
			t.ChangeType = Remove
		}
	}
}

func (a *AzureADFConfig) Diff(target *AzureADFConfig) {
	compareFactory(a.Factory, target.Factory)
	compareCredentials(a.Credential, target.Credential)
	compareLinkedServices(a.LinkedService, target.LinkedService)
	compareManagedVirtualNetworks(a.ManagedVirtualNetwork, target.ManagedVirtualNetwork)
	compareManagedPrivateEndpoints(a.ManagedPrivateEndpoint, target.ManagedPrivateEndpoint)
	compareIntegrationRuntimes(a.IntegrationRuntime, target.IntegrationRuntime)
	compareDatasets(a.Dataset, target.Dataset)
	compareTriggers(a.Trigger, target.Trigger)
	comparePipelines(a.Pipeline, target.Pipeline)
}
