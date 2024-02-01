package pony

import (
	"os"

	"github.com/gofrontier-com/go-utils/output"
	"github.com/gofrontier-com/pony-express/pkg/cmd/cli/apply"
	"github.com/gofrontier-com/pony-express/pkg/cmd/cli/validate"
	vers "github.com/gofrontier-com/pony-express/pkg/cmd/cli/version"
	"github.com/gofrontier-com/pony-express/pkg/util/app_config"
	"github.com/spf13/cobra"
)

func NewRootCmd(version string, commit string, date string) *cobra.Command {
	_, err := app_config.LoadAppConfig()
	if err != nil {
		output.PrintlnError(err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:                   "pony",
		DisableFlagsInUseLine: true,
		Short:                 "pony is the command line tool for pony-express",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.AddCommand(apply.NewCmdApply())
	rootCmd.AddCommand(validate.NewCmdValidate())
	rootCmd.AddCommand(vers.NewCmdVersion(version, commit, date))

	return rootCmd
}
