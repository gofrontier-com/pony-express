package adf

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch/v5"
	json "github.com/json-iterator/go"

	"gopkg.in/yaml.v2"
)

type Change struct {
	Name  string
	Type  string
	Patch string
}

func (c *Change) addJsonPatch(patch interface{}) error {
	js, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	c.Patch = string(js)
	return nil
}

func getJsonPatches(adfChanges map[string]interface{}) ([]*Change, error) {
	var changes []*Change
	for k, chg := range adfChanges {
		changesBytes, err := yaml.Marshal(chg)
		if err != nil {
			return nil, err
		}

		allChanges := make([]map[string]interface{}, 0)
		yaml.Unmarshal(changesBytes, &allChanges)

		for _, chg := range allChanges {
			name, ok := chg["name"].(string)
			if !ok {
				name = ""
			}
			c := Change{Name: name, Type: k}
			c.addJsonPatch(chg["patch"])
			changes = append(changes, &c)
		}
	}
	return changes, nil
}

func processDatasetChange(c *Change, datasets []*PonyDataset) error {
	for _, p := range datasets {
		if *p.Dataset.Name == c.Name {
			ij, err := p.Dataset.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			p.Dataset.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processLinkedServiceChange(c *Change, linkedServices []*PonyLinkedService) error {
	for _, p := range linkedServices {
		if *p.LinkedService.Name == c.Name {
			ij, err := p.LinkedService.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			p.LinkedService.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processPipelineChange(c *Change, pipelines []*PonyPipeline) error {
	for _, p := range pipelines {
		if *p.Pipeline.Name == c.Name {
			ij, err := p.Pipeline.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			p.Pipeline.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processIntegrationRuntimeChange(c *Change, integrationRuntimes []*PonyIntegrationRuntime) error {
	for _, ir := range integrationRuntimes {
		if *ir.IntegrationRuntime.Name == c.Name {
			ij, err := ir.IntegrationRuntime.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			ir.IntegrationRuntime.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processManagedVirtualNetworkChange(c *Change, managedVirtualNetworks []*PonyManagedVirtualNetwork) error {
	for _, mvn := range managedVirtualNetworks {
		if *mvn.ManagedVirtualNetwork.Name == c.Name {
			ij, err := mvn.ManagedVirtualNetwork.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			mvn.ManagedVirtualNetwork.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processManagedPrivateEndpointChange(c *Change, managedPrivateEndpoints []*PonyManagedPrivateEndpoint) error {
	for _, mpe := range managedPrivateEndpoints {
		if *mpe.ManagedPrivateEndpoint.Name == c.Name {
			ij, err := mpe.ManagedPrivateEndpoint.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			mpe.ManagedPrivateEndpoint.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processTriggerChange(c *Change, triggers []*PonyTrigger) error {
	for _, t := range triggers {
		if *t.Trigger.Name == c.Name {
			ij, err := t.Trigger.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			t.Trigger.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func processFactoryChange(c *Change, factory *PonyFactory) error {
	fj, err := factory.Factory.MarshalJSON()
	if err != nil {
		return err
	}
	modifiedAlternative, err := jsonpatch.MergePatch(fj, []byte(c.Patch))
	if err != nil {
		return err
	}
	factory.Factory.UnmarshalJSON(modifiedAlternative)
	return nil
}

func processCredentialChange(c *Change, creds []*PonyCredential) error {
	for _, cr := range creds {
		if *cr.Credential.Name == c.Name {
			ij, err := cr.Credential.MarshalJSON()
			if err != nil {
				return err
			}
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			cr.Credential.UnmarshalJSON(modifiedAlternative)
		}
	}
	return nil
}

func (a *AzureADFConfig) processChanges(changes []*Change) {
	for _, c := range changes {
		switch c.Type {
		case "pipeline":
			processPipelineChange(c, a.Pipeline)
		case "dataset":
			processDatasetChange(c, a.Dataset)
		case "linkedService":
			processLinkedServiceChange(c, a.LinkedService)
		case "integrationRuntime":
			processIntegrationRuntimeChange(c, a.IntegrationRuntime)
		case "managedVirtualNetwork":
			processManagedVirtualNetworkChange(c, a.ManagedVirtualNetwork)
		case "managedPrivateEndpoint":
			processManagedPrivateEndpointChange(c, a.ManagedPrivateEndpoint)
		case "factory":
			processFactoryChange(c, a.Factory)
		case "trigger":
			processTriggerChange(c, a.Trigger)
		case "credential":
			processCredentialChange(c, a.Credential)
		default:
			fmt.Println("Unknown change type: ", c.Type)
		}
	}
}

func (a *AzureADFConfig) ProcessChanges(adfChanges map[string]interface{}) error {
	changes, err := getJsonPatches(adfChanges)
	if err != nil {
		return err
	}

	a.processChanges(changes)
	return nil
}
