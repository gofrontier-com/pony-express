package adf

func (p *PonyManagedPrivateEndpoint) AddDependency(pipeline PonyResource) {
}

func (p *PonyManagedPrivateEndpoint) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyManagedPrivateEndpoint) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyManagedPrivateEndpoint) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyManagedPrivateEndpoint) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyManagedPrivateEndpoint) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyManagedPrivateEndpoint) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyManagedPrivateEndpoint) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyManagedPrivateEndpoint) GetName() *string {
	return p.ManagedPrivateEndpoint.Name
}

func (p *PonyManagedPrivateEndpoint) ToJSON() []byte {
	bytes, _ := p.ManagedPrivateEndpoint.MarshalJSON()
	return bytes
}

func (p *PonyManagedPrivateEndpoint) FromJSON(bytes []byte) {
	p.ManagedPrivateEndpoint.UnmarshalJSON(bytes)
}
