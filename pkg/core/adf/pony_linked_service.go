package adf

func (p *PonyLinkedService) AddDependency(pipeline PonyResource) {
}

func (p *PonyLinkedService) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyLinkedService) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyLinkedService) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyLinkedService) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyLinkedService) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyLinkedService) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyLinkedService) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyLinkedService) GetName() *string {
	return p.LinkedService.Name
}

func (p *PonyLinkedService) ToJSON() []byte {
	bytes, _ := p.LinkedService.MarshalJSON()
	return bytes
}

func (p *PonyLinkedService) FromJSON(bytes []byte) {
	p.LinkedService.UnmarshalJSON(bytes)
}
