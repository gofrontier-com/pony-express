package adf

import (
	"fmt"
	"strings"

	"github.com/go-test/deep"
)

func filterRemoteProps(prop string, extraProps ...string) bool {
	props := []string{"Etag", "ID", "Type", "ConfiguredForDeployment", "RequiresDeployment", "ChangeType"}
	if extraProps != nil {
		props = append(props, extraProps...)
	}
	for _, p := range props {
		if prop == p {
			return false
		}
	}
	return true
}

func compareFactory(source PonyResource, target PonyResource) {
	if diff := deep.Equal(source, target); diff != nil {
		for _, d := range diff {
			prop := strings.Split(d, ":")[0]
			if filterRemoteProps(prop) {
				source.SetRequiresDeployment(true)
				source.SetChangeType(Update)
			}
		}
	}

}

func findMatchingTarget(sourceName *string, target []PonyResource) (PonyResource, error) {
	for _, t := range target {
		if *sourceName == *t.GetName() {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no matching target found for %s", *sourceName)
}

func compare(source []PonyResource, target []PonyResource) {
	for _, s := range source {
		s.SetRequiresDeployment(false)
		t, err := findMatchingTarget(s.GetName(), target)
		if err != nil {
			s.SetRequiresDeployment(true)
			s.SetChangeType(Add)
			continue
		}

		if diff := deep.Equal(s, t); diff != nil {
			for _, d := range diff {
				prop := strings.Split(d, ":")[0]
				if filterRemoteProps(prop) {
					s.SetRequiresDeployment(true)
					s.SetChangeType(Update)
				}
			}
		}
	}

	for _, t := range target {
		_, err := findMatchingTarget(t.GetName(), source)
		if err != nil {
			t.SetRequiresDeployment(true)
			t.SetChangeType(Remove)
		}
	}
}

func (a *PonyADF) Diff(target *PonyADF) {
	compareFactory(a.Factory, target.Factory)
	compare(a.Credential, target.Credential)
	compare(a.LinkedService, target.LinkedService)
	compare(a.ManagedVirtualNetwork, target.ManagedVirtualNetwork)
	compare(a.ManagedPrivateEndpoint, target.ManagedPrivateEndpoint)
	compare(a.IntegrationRuntime, target.IntegrationRuntime)
	compare(a.Dataset, target.Dataset)
	compare(a.Trigger, target.Trigger)
	compare(a.Pipeline, target.Pipeline)
}
