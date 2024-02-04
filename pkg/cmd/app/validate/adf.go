package validate

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gofrontier-com/go-utils/output"
	adfutil "github.com/gofrontier-com/pony-express/pkg/util/adf"
)

func ValidateADF(adfDir string, configFile string, subscriptionid string, resourceGroup string, factoryName string) error {
	myFigure := figure.NewFigure("Pony Express", "doom", true)
	myFigure.Print()

	output.PrintlnfInfo("Loading and validating ADF source from %s", adfDir)
	output.PrintlnfInfo("Loading and validating ADF config from %s\n", configFile)

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

	output.PrintlnfInfo("Valid")

	return nil
}
