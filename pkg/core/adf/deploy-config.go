package adf

func hasWildcard(strArr []string) bool {
	for _, s := range strArr {
		if s == "*" {
			return true
		}
	}
	return false
}

func setDeploymentConfig(config []string, resources []PonyResource) {
	hw := hasWildcard(config)
	for _, resource := range resources {
		if hw {
			resource.SetConfiguredForDeployment(true)
			continue
		}
		for _, cfg := range config {
			if *resource.GetName() == cfg {
				resource.SetConfiguredForDeployment(true)
			}
		}
	}
}

func setTargetDeploymentConfig(config []string, resources []PonyResource) {
	hw := hasWildcard(config)
	if hw {
		return
	}
	for _, resource := range resources {
		match := false
		for _, configPipeline := range config {
			if *resource.GetName() == configPipeline {
				match = true
			}
		}
		if !match {
			resource.SetConfiguredForDeployment(true)
			resource.SetChangeType(Remove)
		}
	}
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
