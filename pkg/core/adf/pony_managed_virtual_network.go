package adf

func (p *PonyManagedVirtualNetwork) AddDependency(pipeline PonyResource) {
}

func (p *PonyManagedVirtualNetwork) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyManagedVirtualNetwork) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyManagedVirtualNetwork) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyManagedVirtualNetwork) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyManagedVirtualNetwork) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyManagedVirtualNetwork) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyManagedVirtualNetwork) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyManagedVirtualNetwork) GetName() *string {
	return p.ManagedVirtualNetwork.Name
}

func (p *PonyManagedVirtualNetwork) ToJSON() []byte {
	bytes, _ := p.ManagedVirtualNetwork.MarshalJSON()
	return bytes
}

func (p *PonyManagedVirtualNetwork) FromJSON(bytes []byte) {
	p.ManagedVirtualNetwork.UnmarshalJSON(bytes)
}
