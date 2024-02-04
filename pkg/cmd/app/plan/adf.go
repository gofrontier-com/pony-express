package plan

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/gofrontier-com/go-utils/output"
	"github.com/gofrontier-com/pony-express/pkg/core/adf"
	adfutil "github.com/gofrontier-com/pony-express/pkg/util/adf"
)

func printToPrint(header string, toPrint []string) {
	if len(toPrint) > 0 {
		output.PrintlnfInfo(header)
		for _, i := range toPrint {
			output.PrintfInfo(i)
		}
		output.PrintlnfInfo("")
	}
}

func printRemovePlan(targetAdf *adf.AzureADFConfig) {
	toPrint := make([]string, 0)
	printed := false

	for _, i := range targetAdf.Credential {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Credential.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Credentials:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.IntegrationRuntime {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.IntegrationRuntime.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Integration Runtimes:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.LinkedService {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.LinkedService.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Linked Services:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.ManagedPrivateEndpoint {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.ManagedPrivateEndpoint.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Managed Private Endpoints:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.ManagedVirtualNetwork {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.ManagedVirtualNetwork.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Managed Virtual Networks:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.Dataset {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Dataset.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Datasets:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.Trigger {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Trigger.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Triggers:", toPrint)
	toPrint = nil

	for _, i := range targetAdf.Pipeline {
		if i.ConfiguredForDeployment && i.ChangeType == adf.Remove {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Pipeline.Name))
			for _, j := range i.Dependencies {
				toPrint = append(toPrint, fmt.Sprintf("  - %s\n", *j.Name))
			}
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Pipelines:", toPrint)
	toPrint = nil

	if !printed {
		output.PrintlnfInfo("(none)")
	}
}

func printPlan(changeType int, sourceAdf *adf.AzureADFConfig) {
	toPrint := make([]string, 0)
	printed := false

	for _, i := range sourceAdf.Credential {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Credential.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Credentials:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.IntegrationRuntime {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.IntegrationRuntime.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Integration Runtimes:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.LinkedService {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.LinkedService.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Linked Services:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.ManagedPrivateEndpoint {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.ManagedPrivateEndpoint.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Managed Private Endpoints:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.ManagedVirtualNetwork {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.ManagedVirtualNetwork.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Managed Virtual Networks:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.Dataset {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Dataset.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Datasets:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.Trigger {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Trigger.Name))
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Triggers:", toPrint)
	toPrint = nil

	for _, i := range sourceAdf.Pipeline {
		if i.RequiresDeployment && i.ConfiguredForDeployment && i.ChangeType == changeType {
			toPrint = append(toPrint, fmt.Sprintf("- %s\n", *i.Pipeline.Name))
			for _, j := range i.Dependencies {
				toPrint = append(toPrint, fmt.Sprintf("  - %s\n", *j.Name))
			}
		}
	}
	if !printed && len(toPrint) > 0 {
		printed = true
	}
	printToPrint("Pipelines:", toPrint)
	toPrint = nil

	if !printed {
		output.PrintlnfInfo("(none)")
	}
}

func PlanADF(adfDir string, configFile string, subscriptionid string, resourceGroup string, factoryName string) error {
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

	targetAdf, err := adf.Fetch(subscriptionid, resourceGroup, factoryName)
	if err != nil {
		return err
	}

	targetAdf.SetTargetDeploymentConfig(&cfg.Deploy)

	sourceAdf.Diff(targetAdf)

	err = sourceAdf.Deps()
	if err != nil {
		return err
	}

	output.PrintlnfInfo("Pony is ready to go!")

	output.PrintlnfInfo("Pony will perform the following actions:")

	output.PrintlnfInfo("\nAdd\n===\n")
	printPlan(adf.Add, sourceAdf)

	output.PrintlnfInfo("\nUpdate\n======\n")
	printPlan(adf.Update, sourceAdf)

	output.PrintlnfInfo("\nRemove\n======\n")
	printRemovePlan(targetAdf)
	return nil
}
