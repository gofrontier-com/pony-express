package adf

// PonyResource is an interface for all ADF resources
type PonyResource interface {
	Base() interface{}
	CheckDependencies() bool
	GetChangeType() int
	GetConfiguredForDeployment() bool
	GetDependencies(...[]PonyResource) []PonyResource
	GetName() *string
	GetRequiresDeployment() bool
	SetChangeType(int)
	SetConfiguredForDeployment(bool)
	SetRequiresDeployment(bool)

	ToJSON() []byte
	FromJSON([]byte)
}
