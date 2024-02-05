package adf

func (p *PonyFactory) AddDependency(pipeline PonyResource) {
}

func (p *PonyFactory) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyFactory) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyFactory) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyFactory) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyFactory) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyFactory) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyFactory) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyFactory) GetName() *string {
	return p.Factory.Name
}

func (p *PonyFactory) ToJSON() []byte {
	bytes, _ := p.Factory.MarshalJSON()
	return bytes
}

func (p *PonyFactory) FromJSON(bytes []byte) {
	p.Factory.UnmarshalJSON(bytes)
}
