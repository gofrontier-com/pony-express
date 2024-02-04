package plan

import (
	"os"
	"path/filepath"

	"github.com/gofrontier-com/pony-express/pkg/cmd/app/plan"
	"github.com/spf13/cobra"
)

var (
	adfDir         string
	configFilePath string
	subscriptionId string
	resourceGroup  string
	factoryName    string
)

// NewCmdVPlan creates a command to plan the Azure Rm config
func NewCmdPlanADF() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adf",
		Short: "Plan ADF",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := plan.PlanADF(adfDir, configFilePath, subscriptionId, resourceGroup, factoryName); err != nil {
				return err
			}

			return nil
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd.Flags().StringVarP(&adfDir, "adf-dir", "a", wd, "ADF source directory")
	cmd.Flags().StringVarP(&configFilePath, "config", "c", filepath.Join(wd, "config.yaml"), "Config file")
	cmd.Flags().StringVarP(&subscriptionId, "subscription-id", "s", "", "Subscription ID")
	cmd.Flags().StringVarP(&resourceGroup, "resource-group", "g", "", "Resource group")
	cmd.Flags().StringVarP(&factoryName, "adf-name", "n", "", "ADF name")

	return cmd
}
