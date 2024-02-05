package adf

func (p *PonyDataset) AddDependency(pipeline PonyResource) {
}

func (p *PonyDataset) GetDependencies(resource []PonyResource) []PonyResource {
	return nil
}

func (p *PonyDataset) SetConfiguredForDeployment(d bool) {
	p.ConfiguredForDeployment = d
}

func (p *PonyDataset) SetRequiresDeployment(d bool) {
	p.RequiresDeployment = d
}

func (p *PonyDataset) SetChangeType(ct int) {
	p.ChangeType = ct
}

func (p *PonyDataset) GetChangeType() int {
	return p.ChangeType
}

func (p *PonyDataset) GetConfiguredForDeployment() bool {
	return p.ConfiguredForDeployment
}

func (p *PonyDataset) GetRequiresDeployment() bool {
	return p.RequiresDeployment
}

func (p *PonyDataset) GetName() *string {
	return p.Dataset.Name
}

func (p *PonyDataset) ToJSON() []byte {
	bytes, _ := p.Dataset.MarshalJSON()
	return bytes
}

func (p *PonyDataset) FromJSON(bytes []byte) {
	p.Dataset.UnmarshalJSON(bytes)
}
