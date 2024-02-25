package plan

import (
	"github.com/spf13/cobra"
)

// NewCmdPlan creates a command to Plan config
func NewCmdPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "plan",
	}

	cmd.AddCommand(NewCmdPlanADF())

	return cmd
}
