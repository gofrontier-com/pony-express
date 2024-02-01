package adf

func hasWildcard(strArr []string) bool {
	for _, s := range strArr {
		if s == "*" {
			return true
		}
	}
	return false
}

func setDeploymentConfigCredentials(config []string, credentials []*PonyCredential) {
	hw := hasWildcard(config)
	for _, credential := range credentials {
		if hw {
			credential.ConfiguredForDeployment = true
			continue
		}
		for _, configCredential := range config {
			if *credential.Credential.Name == configCredential {
				credential.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigDatasets(config []string, datasets []*PonyDataset) {
	hw := hasWildcard(config)
	for _, dataset := range datasets {
		if hw {
			dataset.ConfiguredForDeployment = true
			continue
		}
		for _, configDataset := range config {
			if *dataset.Dataset.Name == configDataset {
				dataset.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigIntegrationRuntimes(config []string, integrationRuntimes []*PonyIntegrationRuntime) {
	hw := hasWildcard(config)
	for _, integrationRuntime := range integrationRuntimes {
		if hw {
			integrationRuntime.ConfiguredForDeployment = true
			continue
		}
		for _, configIntegrationRuntime := range config {
			if *integrationRuntime.IntegrationRuntime.Name == configIntegrationRuntime {
				integrationRuntime.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigLinkedServices(config []string, linkedServices []*PonyLinkedService) {
	hw := hasWildcard(config)
	for _, linkedService := range linkedServices {
		if hw {
			linkedService.ConfiguredForDeployment = true
			continue
		}
		for _, configLinkedService := range config {
			if *linkedService.LinkedService.Name == configLinkedService {
				linkedService.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigManagedPrivateEndpoints(config []string, managedPrivateEndpoints []*PonyManagedPrivateEndpoint) {
	hw := hasWildcard(config)
	for _, managedPrivateEndpoint := range managedPrivateEndpoints {
		if hw {
			managedPrivateEndpoint.ConfiguredForDeployment = true
			continue
		}
		for _, configManagedPrivateEndpoint := range config {
			if *managedPrivateEndpoint.ManagedPrivateEndpoint.Name == configManagedPrivateEndpoint {
				managedPrivateEndpoint.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigManagedVirtualNetworks(config []string, managedVirtualNetworks []*PonyManagedVirtualNetwork) {
	hw := hasWildcard(config)
	for _, managedVirtualNetwork := range managedVirtualNetworks {
		if hw {
			managedVirtualNetwork.ConfiguredForDeployment = true
			continue
		}
		for _, configManagedVirtualNetwork := range config {
			if *managedVirtualNetwork.ManagedVirtualNetwork.Name == configManagedVirtualNetwork {
				managedVirtualNetwork.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigPipelines(config []string, pipelines []*PonyPipeline) {
	hw := hasWildcard(config)
	for _, pipeline := range pipelines {
		if hw {
			pipeline.ConfiguredForDeployment = true
			continue
		}
		for _, configPipeline := range config {
			if *pipeline.Pipeline.Name == configPipeline {
				pipeline.ConfiguredForDeployment = true
			}
		}
	}
}

func setDeploymentConfigTriggers(config []string, triggers []*PonyTrigger) {
	hw := hasWildcard(config)
	for _, trigger := range triggers {
		if hw {
			trigger.ConfiguredForDeployment = true
			continue
		}
		for _, configTrigger := range config {
			if *trigger.Trigger.Name == configTrigger {
				trigger.ConfiguredForDeployment = true
			}
		}
	}
}

func (a *AzureADFConfig) SetDeploymentConfig(config *ADFDeployConfig) {
	setDeploymentConfigCredentials(config.Credential, a.Credential)

	setDeploymentConfigPipelines(config.Pipeline, a.Pipeline)

	setDeploymentConfigTriggers(config.Trigger, a.Trigger)

	setDeploymentConfigDatasets(config.Dataset, a.Dataset)

	setDeploymentConfigIntegrationRuntimes(config.IntegrationRuntime, a.IntegrationRuntime)

	setDeploymentConfigLinkedServices(config.LinkedService, a.LinkedService)

	setDeploymentConfigManagedPrivateEndpoints(config.ManagedPrivateEndpoint, a.ManagedPrivateEndpoint)

	setDeploymentConfigManagedVirtualNetworks(config.ManagedVirtualNetwork, a.ManagedVirtualNetwork)
}
