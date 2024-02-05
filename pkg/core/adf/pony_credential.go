package adf

func (p *PonyCredential) AddDependency(pipeline PonyResource) {
}

func (p *PonyCredential) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyCredential) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyCredential) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyCredential) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyCredential) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyCredential) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyCredential) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyCredential) GetName() *string {
	return p.Credential.Name
}

func (p *PonyCredential) ToJSON() []byte {
	bytes, _ := p.Credential.MarshalJSON()
	return bytes
}

func (p *PonyCredential) FromJSON(bytes []byte) {
	p.Credential.UnmarshalJSON(bytes)
}
