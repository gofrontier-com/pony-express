package adf

import (
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

func (p *PonyLinkedService) getDependantCredentials(credentials []PonyResource) {
	props := p.LinkedService.Properties
	lsType := p.LinkedService.Properties.GetLinkedService().Type
	var lsName *string

	switch *lsType {
	case "AzureBlobFS":
		pp := props.(*armdatafactory.AzureBlobFSLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	case "AzureBlobStorage":
		pp := props.(*armdatafactory.AzureBlobStorageLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	case "AzureDataLakeStore":
		pp := props.(*armdatafactory.AzureDataLakeStoreLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	// case "AzureFileStorage":
	// 	pp := props.(*armdatafactory.AzureFileStorageLinkedService)
	// 	lsName = pp.TypeProperties.Credential.ReferenceName
	case "AzureKeyVault":
		pp := props.(*armdatafactory.AzureKeyVaultLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	case "AzureSqlDatabase":
		pp := props.(*armdatafactory.AzureSQLDatabaseLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	case "AzureSQLDW":
		pp := props.(*armdatafactory.AzureSQLDWLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	case "CosmosDb":
		pp := props.(*armdatafactory.CosmosDbLinkedService)
		lsName = pp.TypeProperties.Credential.ReferenceName
	default:
		fmt.Println("Unknown linked service type: ", *lsType)
		return
	}

	for _, cred := range credentials {
		if *cred.GetName() == *lsName {
			p.AddDependency(cred)
		}
	}
}

func (p *PonyLinkedService) getDependantIntegrationRuntimes(integrationRuntimes []PonyResource) {
	props := p.LinkedService.Properties
	cv := props.GetLinkedService().ConnectVia
	if cv == nil {
		return
	}
	irName := cv.ReferenceName

	for _, ir := range integrationRuntimes {
		if *ir.GetName() == *irName {
			p.AddDependency(ir)
		}
	}
}

func (p *PonyLinkedService) AddDependency(pipeline PonyResource) {
	p.Dependencies = append(p.Dependencies, pipeline)
}

func (p *PonyLinkedService) Base() interface{} {
	return p.LinkedService
}

func (p *PonyLinkedService) CheckDependencies() bool {
	if !p.ConfiguredForDeployment {
		return true
	}

	for _, dep := range p.Dependencies {
		if !dep.GetConfiguredForDeployment() {
			fmt.Println("LinkedService ", *p.GetName(), " has a dependency that is not configured for deployment: ", *dep.GetName())
			return false
		}
	}

	return true
}

func (p *PonyLinkedService) GetDependencies(resources ...[]PonyResource) []PonyResource {
	var credentials []PonyResource
	var integrationRuntimes []PonyResource

	for i, resource := range resources {
		for _, r := range resource {
			t := reflect.TypeOf(r).Elem().Name()
			if t == "PonyCredential" {
				credentials = resources[i]
			} else if t == "PonyIntegrationRuntime" {
				integrationRuntimes = resources[i]
			}
			break
		}
	}

	p.getDependantCredentials(credentials)
	p.getDependantIntegrationRuntimes(integrationRuntimes)

	return p.Dependencies
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
