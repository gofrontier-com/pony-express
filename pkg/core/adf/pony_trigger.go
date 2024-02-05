package adf

func (p *PonyTrigger) AddDependency(pipeline PonyResource) {
}

func (p *PonyTrigger) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyTrigger) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyTrigger) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyTrigger) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyTrigger) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyTrigger) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyTrigger) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyTrigger) GetName() *string {
	return p.Trigger.Name
}

func (p *PonyTrigger) ToJSON() []byte {
	bytes, _ := p.Trigger.MarshalJSON()
	return bytes
}

func (p *PonyTrigger) FromJSON(bytes []byte) {
	p.Trigger.UnmarshalJSON(bytes)
}
