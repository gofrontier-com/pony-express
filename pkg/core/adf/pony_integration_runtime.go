package adf

func (p *PonyIntegrationRuntime) AddDependency(pipeline PonyResource) {
}

func (p *PonyIntegrationRuntime) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyIntegrationRuntime) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyIntegrationRuntime) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyIntegrationRuntime) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyIntegrationRuntime) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyIntegrationRuntime) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyIntegrationRuntime) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyIntegrationRuntime) GetName() *string {
	return p.IntegrationRuntime.Name
}

func (p *PonyIntegrationRuntime) ToJSON() []byte {
	bytes, _ := p.IntegrationRuntime.MarshalJSON()
	return bytes
}

func (p *PonyIntegrationRuntime) FromJSON(bytes []byte) {
	p.IntegrationRuntime.UnmarshalJSON(bytes)
}
