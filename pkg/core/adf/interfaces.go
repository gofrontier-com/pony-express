package adf

type PonyResource interface {
	GetName() *string
	GetConfiguredForDeployment() bool
	GetChangeType() int
	GetDependencies() []PonyResource
	getPipelineDeps([]PonyResource) error
	GetRequiresDeployment() bool
	SetRequiresDeployment(bool)
	SetChangeType(int)
	SetConfiguredForDeployment(bool)

	ToJSON() []byte
	FromJSON([]byte)
}
