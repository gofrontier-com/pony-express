package plan

import (
	"fmt"
	"sync"

	"github.com/common-nighthawk/go-figure"
	"github.com/gofrontier-com/go-utils/output"
	"github.com/gofrontier-com/pony-express/pkg/core/adf"
	adfutil "github.com/gofrontier-com/pony-express/pkg/util/adf"
)

func printToPrint(header string, toPrint []string) bool {
	if len(toPrint) > 0 {
		output.PrintlnfInfo(header)
		for _, i := range toPrint {
			output.PrintfInfo(i)
		}
		output.PrintlnfInfo("")
		return true
	}
	return false
}

func printResourcePlan(header string, resource []adf.PonyResource, changeType int) bool {
	toPrint := make([]string, 0)

	for _, r := range resource {
		condition := r.GetRequiresDeployment() && r.GetConfiguredForDeployment()
		if changeType != adf.Remove {
			condition = condition && r.GetChangeType() == changeType
		}
		if condition {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *r.GetName()))
		}
	}

	return printToPrint(header, toPrint)
}

func printPlan(a *adf.PonyADF, changeType int) {
	printed := false

	printed = printResourcePlan("Credentials:", a.Credential, changeType) || printed
	printed = printResourcePlan("Integration Runtimes:", a.IntegrationRuntime, changeType) || printed
	printed = printResourcePlan("Linked Services:", a.LinkedService, changeType) || printed
	printed = printResourcePlan("Managed Private Endpoints:", a.ManagedPrivateEndpoint, changeType) || printed
	printed = printResourcePlan("Managed Virtual Networks:", a.ManagedVirtualNetwork, changeType) || printed
	printed = printResourcePlan("Datasets:", a.Dataset, changeType) || printed
	printed = printResourcePlan("Triggers:", a.Trigger, changeType) || printed
	printed = printResourcePlan("Pipeline:", a.Pipeline, changeType) || printed

	if !printed {
		output.PrintlnfInfo("(none)")
	}
}

func doSource(wg *sync.WaitGroup, sourceAdf *adf.PonyADF, cfg *adf.PonyConfig, adfDir string) {
	defer wg.Done()

	sourceAdf.LoadFromFolder(adfDir)

	sourceAdf.ProcessChanges(cfg.Changes)

	sourceAdf.SetDeploymentConfig(&cfg.Deploy)
}

func doTarget(wg *sync.WaitGroup, targetAdf *adf.PonyADF, cfg *adf.PonyConfig) {
	defer wg.Done()

	targetAdf.Fetch()

	targetAdf.SetTargetDeploymentConfig(&cfg.Deploy)
}

func PlanADF(adfDir string, configFile string, subscriptionid string, resourceGroup string, factoryName string) error {
	myFigure := figure.NewFigure("Pony Express", "doom", true)
	myFigure.Print()

	output.PrintlnfInfo("Loading and validating ADF source from %s", adfDir)
	output.PrintlnfInfo("Loading and validating ADF config from %s\n", configFile)

	cfg, err := adfutil.LoadConfig(configFile)
	if err != nil {
		return err
	}

	sourceAdf := adf.NewADF()
	targetAdf, _ := adf.NewRemoteADF(subscriptionid, resourceGroup, factoryName)

	var wg sync.WaitGroup

	wg.Add(2)

	go doSource(&wg, sourceAdf, cfg, adfDir)
	go doTarget(&wg, targetAdf, cfg)

	wg.Wait()

	sourceAdf.Diff(targetAdf)

	err = sourceAdf.Deps()
	if err != nil {
		return err
	}

	output.PrintlnfInfo("Pony is ready to go!")

	output.PrintlnfInfo("Pony will perform the following actions:")

	output.PrintlnfInfo("\nAdd\n===\n")
	printPlan(sourceAdf, adf.Add)

	output.PrintlnfInfo("\nUpdate\n======\n")
	printPlan(sourceAdf, adf.Update)

	output.PrintlnfInfo("\nRemove\n======\n")
	printPlan(targetAdf, adf.Remove)

	return nil
}
