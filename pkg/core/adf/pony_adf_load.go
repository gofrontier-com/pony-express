package adf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// LoadFromFolder loads the ADF resources from the specified folder.
func (a *PonyADF) LoadFromFolder(configDir string) error {
	err := filepath.Walk(configDir, func(path string, f os.FileInfo, err error) error {
		s := strings.Replace(path, configDir+string(filepath.Separator), "", 1)
		ss := strings.Split(s, string(filepath.Separator))

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			if len(ss) > 1 {
				property := ss[len(ss)-2]
				switch property {
				case "factory":
					err = a.LoadFactory(path)
				case "credential":
					err = a.LoadCredential(path)
				case "linkedService":
					err = a.LoadLinkedService(path)
				case "managedVirtualNetwork":
					err = a.LoadManagedVirtualNetwork(path)
				case "managedPrivateEndpoint":
					err = a.LoadManagedPrivateEndPoint(path)
				case "integrationRuntime":
					err = a.LoadIntegrationRuntime(path)
				case "dataset":
					err = a.LoadDataset(path)
				case "trigger":
					err = a.LoadTrigger(path)
				case "pipeline":
					err = a.LoadPipeline(path)
				default:
					fmt.Println("Not implemented or not used: ", property)
				}
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	return err
}

// LoadCredential loads the credential from the specified file.
func (a *PonyADF) LoadCredential(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	cred := &armdatafactory.ManagedIdentityCredentialResource{}
	cred.UnmarshalJSON(b)
	c := &PonyCredential{
		Credential: cred,
	}
	a.Credential = append(a.Credential, c)
	return nil
}

// LoadDataset loads the dataset from the specified file.
func (a *PonyADF) LoadDataset(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	dsr := &armdatafactory.DatasetResource{}
	dsr.UnmarshalJSON(b)
	ds := &PonyDataset{
		Dataset: dsr,
	}
	a.Dataset = append(a.Dataset, ds)
	return nil
}

// LoadFactory loads the factory from the specified file.
func (a *PonyADF) LoadFactory(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	af := &armdatafactory.Factory{}
	af.UnmarshalJSON(b)
	f := &PonyFactory{
		Factory:                 af,
		ConfiguredForDeployment: true,
	}
	a.Factory = f
	return nil
}

// LoadIntegrationRuntime loads the integration runtime from the specified file.
func (a *PonyADF) LoadIntegrationRuntime(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	irr := &armdatafactory.IntegrationRuntimeResource{}
	irr.UnmarshalJSON(b)
	ir := &PonyIntegrationRuntime{
		IntegrationRuntime: irr,
	}
	a.IntegrationRuntime = append(a.IntegrationRuntime, ir)
	return nil
}

// LoadLinkedService loads the linked service from the specified file.
func (a *PonyADF) LoadLinkedService(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	lsr := &armdatafactory.LinkedServiceResource{}
	lsr.UnmarshalJSON(b)
	ls := &PonyLinkedService{
		LinkedService: lsr,
	}
	a.LinkedService = append(a.LinkedService, ls)
	return nil
}

// LoadManagedPrivateEndPoint loads the managed private endpoint from the specified file.
func (a *PonyADF) LoadManagedPrivateEndPoint(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	mper := &armdatafactory.ManagedPrivateEndpointResource{}
	mper.UnmarshalJSON(b)
	mpe := &PonyManagedPrivateEndpoint{
		ManagedPrivateEndpoint: mper,
	}
	a.ManagedPrivateEndpoint = append(a.ManagedPrivateEndpoint, mpe)
	return nil
}

// LoadManagedVirtualNetwork loads the managed virtual network from the specified file.
func (a *PonyADF) LoadManagedVirtualNetwork(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	mvnr := &armdatafactory.ManagedVirtualNetworkResource{}
	mvnr.UnmarshalJSON(b)
	mvn := &PonyManagedVirtualNetwork{
		ManagedVirtualNetwork: mvnr,
	}
	a.ManagedVirtualNetwork = append(a.ManagedVirtualNetwork, mvn)
	return nil
}

// LoadPipeline loads the pipeline from the specified file.
func (a *PonyADF) LoadPipeline(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	ls := &armdatafactory.PipelineResource{}
	ls.UnmarshalJSON(b)
	p := &PonyPipeline{
		Pipeline: ls,
	}
	a.Pipeline = append(a.Pipeline, p)
	return nil
}

// LoadTrigger loads the trigger from the specified file.
func (a *PonyADF) LoadTrigger(filePath string) error {
	b, err := getJsonBytes(filePath)
	if err != nil {
		return err
	}

	trigger := &armdatafactory.TriggerResource{}
	trigger.UnmarshalJSON(b)
	t := &PonyTrigger{
		Trigger: trigger,
	}
	a.Trigger = append(a.Trigger, t)
	return nil
}
