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

func processChanges(c *Change, resources []PonyResource) error {
	for _, r := range resources {
		if *r.GetName() == c.Name {
			processChange(c, r)
		}
	}
	return nil
}

func processChange(c *Change, resource PonyResource) error {
	fj := resource.ToJSON()
	modifiedAlternative, err := jsonpatch.MergePatch(fj, []byte(c.Patch))
	if err != nil {
		return err
	}
	resource.FromJSON(modifiedAlternative)
	return nil
}

func (a *PonyADF) processChanges(changes []*Change) {
	for _, c := range changes {
		switch c.Type {
		case "pipeline":
			processChanges(c, a.Pipeline)
		case "dataset":
			processChanges(c, a.Dataset)
		case "linkedService":
			processChanges(c, a.LinkedService)
		case "integrationRuntime":
			processChanges(c, a.IntegrationRuntime)
		case "managedVirtualNetwork":
			processChanges(c, a.ManagedVirtualNetwork)
		case "managedPrivateEndpoint":
			processChanges(c, a.ManagedPrivateEndpoint)
		case "factory":
			processChange(c, a.Factory)
		case "trigger":
			processChanges(c, a.Trigger)
		case "credential":
			processChanges(c, a.Credential)
		default:
			fmt.Println("Unknown change type: ", c.Type)
		}
	}
}

func (a *PonyADF) ProcessChanges(adfChanges map[string]interface{}) error {
	changes, err := getJsonPatches(adfChanges)
	if err != nil {
		return err
	}

	a.processChanges(changes)
	return nil
}
