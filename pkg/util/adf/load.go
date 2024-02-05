package adf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofrontier-com/pony-express/pkg/core/adf"
)

func LoadMap(configDir string, subscriptionId string, resourceGroup string, factoryName string) (*adf.PonyADF, error) {
	configurationData, err := adf.NewADF(subscriptionId, resourceGroup, factoryName)

	err = filepath.Walk(configDir, func(path string, f os.FileInfo, err error) error {
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
					err = configurationData.LoadFactory(path)
					if err != nil {
						return err
					}
				case "credential":
					err = configurationData.LoadCredential(path)
					if err != nil {
						return err
					}
				case "linkedService":
					err = configurationData.LoadLinkedService(path)
					if err != nil {
						return err
					}
				case "managedVirtualNetwork":
					err = configurationData.LoadManagedVirtualNetwork(path)
					if err != nil {
						return err
					}
				case "managedPrivateEndpoint":
					err = configurationData.LoadManagedPrivateEndPoint(path)
					if err != nil {
						return err
					}
				case "integrationRuntime":
					err = configurationData.LoadIntegrationRuntime(path)
					if err != nil {
						return err
					}
				case "dataset":
					err = configurationData.LoadDataset(path)
					if err != nil {
						return err
					}
				case "trigger":
					err = configurationData.LoadTrigger(path)
					if err != nil {
						return err
					}
				case "pipeline":
					err = configurationData.LoadPipeline(path)
					if err != nil {
						return err
					}
				default:
					fmt.Println("Not implemented or not used: ", property)
				}
			}
		}
		return nil
	})

	return configurationData, err
}
