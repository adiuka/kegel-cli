package command

import (
	"fmt"
	"time"

	"kegel-cli/internal/ui"
	"kegel-cli/internal/workout"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:		"start",
	Short:	"Start a Kegel training session",
	Run:	func(cmd *cobra.Command, args []string) {
		plan := workout.Default

		fmt.Printf("Starting session: %d reps, %.0fs squeeze, %.0fs rest\n\n",
			plan.Reps,
			plan.Squeeze.Seconds(),
			plan.Rest.Seconds(),
		)

		for rep := 1; rep <= plan.Reps; rep++ {
			fmt.Printf("	Rep %d/%d\n", rep, plan.Reps)
			runPhase("squeeze", plan.Squeeze)

			runPhase("rest", plan.Rest)

			fmt.Println()
		}

		fmt.Printf("\nDone! Great work.\n")
	},
}

func runPhase(phase string, duration time.Duration) {
	start := time.Now()
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		elapsed := time.Since(start).Seconds()
		ui.RenderBar(phase, elapsed, duration.Seconds())

		if elapsed >= duration.Seconds() {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}