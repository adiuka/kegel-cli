package command

import (
	"fmt"
	"kegel-cli/internal/workout"
	"time"

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

		for rep := 1; rep <= plan.Reps; rep++ {
			fmt.Printf("Rep %d/%d - SQUEEZE\n", rep, plan.Reps)
			time.Sleep(plan.Squeeze)

			fmt.Printf("Rep %d/%d - REST\n", rep, plan.Reps)
			time.Sleep(plan.Rest)
		}

		fmt.Printf("\nDone! Great work.\n")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}