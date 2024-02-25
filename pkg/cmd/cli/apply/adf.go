package apply

import (
	"os"
	"path/filepath"

	"github.com/gofrontier-com/pony-express/pkg/cmd/app/apply"
	"github.com/spf13/cobra"
)

var (
	adfDir         string
	configFilePath string
	subscriptionId string
	resourceGroup  string
	factoryName    string
)

// NewCmdApplyAzureRm creates a command to apply the Azure RM config
func NewCmdApplyADF() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adf",
		Short: "Apply Azure Data Factory",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := apply.ApplyADF(adfDir, configFilePath, subscriptionId, resourceGroup, factoryName); err != nil {
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
