package command

import (
	"fmt"
	"kegel-cli/internal/workout"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:		"start",
	Short:	"Start a Kegel training session",
	Run:	func(cmd *cobra.Command, args []string) {
		plan := workout.Default
		fmt.Printf("Starting session: %d reps, %.0fs squeeze, %.0fs rest\n",
			plan.Reps,
			plan.Squeeze.Seconds(),
			plan.Rest.Seconds(),
		)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}