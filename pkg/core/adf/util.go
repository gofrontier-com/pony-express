package adf

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-test/deep"
	"gopkg.in/yaml.v2"
)

// CreateClientFactory creates a new client factory for the given subscription id.
func CreateClientFactory(subscriptionid string) (*armdatafactory.ClientFactory, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}
	clientFactory, err := armdatafactory.NewClientFactory(subscriptionid, cred, nil)
	if err != nil {
		return nil, err
	}
	return clientFactory, nil
}

func getJsonBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func hasWildcard(strArr []string) bool {
	for _, s := range strArr {
		if s == "*" {
			return true
		}
	}
	return false
}

func setDeploymentConfig(config []string, resources []PonyResource) {
	hw := hasWildcard(config)
	for _, resource := range resources {
		if hw {
			resource.SetConfiguredForDeployment(true)
			continue
		}
		for _, cfg := range config {
			if *resource.GetName() == cfg {
				resource.SetConfiguredForDeployment(true)
			}
		}
	}
}

func setTargetDeploymentConfig(config []string, resources []PonyResource) {
	hw := hasWildcard(config)
	if hw {
		return
	}
	for _, resource := range resources {
		match := false
		for _, configPipeline := range config {
			if *resource.GetName() == configPipeline {
				match = true
			}
		}
		if !match {
			resource.SetConfiguredForDeployment(true)
			resource.SetChangeType(Remove)
		}
	}
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

func filterRemoteProps(prop string, prefix string, extraProps ...string) bool {
	props := []string{
		fmt.Sprintf("%s.Etag", prefix),
		fmt.Sprintf("%s.ID", prefix),
		fmt.Sprintf("%s.Type", prefix),
		"ConfiguredForDeployment",
		"RequiresDeployment",
		"ChangeType"}
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
			if filterRemoteProps(prop, "Factory") {
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

func compare(source []PonyResource, target []PonyResource, prefix string, additionalProps ...string) {
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
				if filterRemoteProps(prop, prefix, additionalProps...) {
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
