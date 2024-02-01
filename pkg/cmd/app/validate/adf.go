package validate

import (
	"fmt"

	"github.com/gofrontier-com/go-utils/output"
	"github.com/gofrontier-com/pony-express/pkg/core/adf"
	adfutil "github.com/gofrontier-com/pony-express/pkg/util/adf"
)

func ValidateADF(adfDir string, configFile string, subscriptionid string, resourceGroup string, factoryName string) error {
	output.PrintlnfInfo("Loading and validating ADF source from %s", adfDir)
	output.PrintlnfInfo("Loading and validating ADF config from %s", configFile)

	sourceAdf, err := adfutil.LoadMap(adfDir, subscriptionid, resourceGroup, factoryName)
	if err != nil {
		return err
	}

	cfg, err := adfutil.LoadConfig(configFile)
	if err != nil {
		return err
	}

	sourceAdf.ProcessChanges(cfg.Changes)

	sourceAdf.SetDeploymentConfig(&cfg.Deploy)

	targetAdf, err := adf.Fetch(subscriptionid, resourceGroup, factoryName)
	if err != nil {
		return err
	}

	sourceAdf.Diff(targetAdf)

	err = sourceAdf.Deps()
	if err != nil {
		return err
	}

	fmt.Println("To deploy pipelines:")
	for _, i := range sourceAdf.Pipeline {
		if i.RequiresDeployment && i.ConfiguredForDeployment {
			fmt.Printf("- %s\n", *i.Pipeline.Name)
			for _, j := range i.Dependencies {
				fmt.Printf("  - %s\n", *j.Name)
			}
		}
	}

	return nil
}
