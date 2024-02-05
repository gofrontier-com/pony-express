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

func processChange(c *Change, items []PonyResource) error {
	for _, i := range items {
		if *i.GetName() == c.Name {
			ij := i.ToJSON()
			modifiedAlternative, err := jsonpatch.MergePatch(ij, []byte(c.Patch))
			if err != nil {
				return err
			}
			i.FromJSON(modifiedAlternative)
		}
	}
	return nil
}

func processFactoryChange(c *Change, factory PonyResource) error {
	fj := factory.ToJSON()
	modifiedAlternative, err := jsonpatch.MergePatch(fj, []byte(c.Patch))
	if err != nil {
		return err
	}
	factory.FromJSON(modifiedAlternative)
	return nil
}

func (a *AzureADFConfig) processChanges(changes []*Change) {
	for _, c := range changes {
		switch c.Type {
		case "pipeline":
			processChange(c, a.Pipeline)
		case "dataset":
			processChange(c, a.Dataset)
		case "linkedService":
			processChange(c, a.LinkedService)
		case "integrationRuntime":
			processChange(c, a.IntegrationRuntime)
		case "managedVirtualNetwork":
			processChange(c, a.ManagedVirtualNetwork)
		case "managedPrivateEndpoint":
			processChange(c, a.ManagedPrivateEndpoint)
		case "factory":
			processFactoryChange(c, a.Factory)
		case "trigger":
			processChange(c, a.Trigger)
		case "credential":
			processChange(c, a.Credential)
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
