package apply

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	jsonpatch "github.com/evanphx/json-patch"
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
		if changeType == adf.Remove {
			condition = r.GetChangeType() == changeType
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
	// printed = printResourcePlan("Managed Virtual Networks:", a.ManagedVirtualNetwork, changeType) || printed
	printed = printResourcePlan("Datasets:", a.Dataset, changeType) || printed
	printed = printResourcePlan("Triggers:", a.Trigger, changeType) || printed
	printed = printResourcePlan("Pipeline:", a.Pipeline, changeType) || printed

	if !printed {
		output.PrintlnfInfo("(none)")
	}
}

func applyRemovals(targetAdf *adf.PonyADF) {
	for _, r := range targetAdf.Trigger {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveTrigger(r)
		}
	}

	for _, r := range targetAdf.Pipeline {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemovePipeline(r)
		}
	}

	for _, r := range targetAdf.Dataset {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveDataset(r)
		}
	}

	for _, r := range targetAdf.LinkedService {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveLinkedService(r)
		}
	}

	for _, r := range targetAdf.Credential {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveCredential(r)
		}
	}

	for _, r := range targetAdf.IntegrationRuntime {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveIntegrationRuntime(r)
		}
	}

	for _, r := range targetAdf.ManagedPrivateEndpoint {
		if r.GetChangeType() == adf.Remove {
			targetAdf.RemoveManagedPrivateEndpoint(r)
		}
	}
}

func applyAdditionsandUpdates(sourceAdf *adf.PonyADF) {
	if (sourceAdf.Factory.GetChangeType() == adf.Add || sourceAdf.Factory.GetChangeType() == adf.Update) && sourceAdf.Factory.GetRequiresDeployment() && sourceAdf.Factory.GetConfiguredForDeployment() {
		sourceAdf.AddOrUpdateFactory()
	}

	for _, r := range sourceAdf.IntegrationRuntime {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateIntegrationRuntime(r)
		}
	}

	for _, r := range sourceAdf.ManagedVirtualNetwork {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateManagedVirtualNetwork(r)
		}
	}

	for _, r := range sourceAdf.ManagedPrivateEndpoint {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateManagedPrivateEndpoint(r)
		}
	}

	for _, r := range sourceAdf.Credential {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateCredential(r)
		}
	}

	for _, r := range sourceAdf.LinkedService {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateLinkedService(r)
		}
	}

	for _, r := range sourceAdf.Dataset {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateDataset(r)
		}
	}

	for _, r := range sourceAdf.Pipeline {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdatePipeline(r)
		}
	}

	for _, r := range sourceAdf.Trigger {
		if (r.GetChangeType() == adf.Add || r.GetChangeType() == adf.Update) && r.GetRequiresDeployment() && r.GetConfiguredForDeployment() {
			sourceAdf.AddOrUpdateTrigger(r)
		}
	}

}

func ApplyADF(adfDir string, configFile string, subscriptionid string, resourceGroup string, factoryName string) error {
	myFigure := figure.NewFigure("Pony Express", "doom", true)
	myFigure.Print()

	output.PrintlnfInfo("Loading and validating ADF source from %s", adfDir)
	output.PrintlnfInfo("Loading and validating ADF config from %s\n", configFile)

	cfg, err := adfutil.LoadConfig(configFile)
	if err != nil {
		return err
	}

	sourceAdf, _ := adf.NewRemoteADF(subscriptionid, resourceGroup, factoryName)
	targetAdf, _ := adf.NewRemoteADF(subscriptionid, resourceGroup, factoryName)

	sourceAdf.LoadFromFolder(adfDir)

	sourceAdf.ProcessChanges(cfg.Changes)

	sourceAdf.SetDeploymentConfig(&cfg.Deploy)

	depsSatisfied := sourceAdf.Deps()
	if depsSatisfied == false {
		return err
	}

	targetAdf.Fetch()

	targetAdf.SetTargetDeploymentConfig(&cfg.Deploy)

	sourceAdf.Diff(targetAdf)

	output.PrintlnfInfo("Pony is ready to go!")

	output.PrintlnfInfo("Pony will perform the following actions:")

	output.PrintlnfInfo("\nAdd\n===\n")
	printPlan(sourceAdf, adf.Add)

	output.PrintlnfInfo("\nUpdate\n======\n")
	printPlan(sourceAdf, adf.Update)

	output.PrintlnfInfo("\nRemove\n======\n")
	printPlan(targetAdf, adf.Remove)

	output.PrintlnfInfo("\nRemove\n======\n")
	applyRemovals(targetAdf)

	output.PrintlnfInfo("\nAdd\n===\n")
	tfj := targetAdf.Factory.ToJSON()
	sfj := sourceAdf.Factory.ToJSON()
	modifiedAlternative, err := jsonpatch.MergePatch(tfj, sfj)
	sourceAdf.Factory.FromJSON(modifiedAlternative)
	applyAdditionsandUpdates(sourceAdf)

	return nil
}
