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

		ui.PrintBanner()
		ui.PrintPlanSummary(plan.Reps, plan.Squeeze, plan.Rest)
		ui.WaitForEnter()

		fmt.Scanln()

		for rep := 1; rep <= plan.Reps; rep++ {
			fmt.Printf("	 Rep %d/%d\n", rep, plan.Reps)
			runPhase("squeeze", plan.Squeeze)

			runPhase("rest", plan.Rest)

			fmt.Println()

			if rep < plan.Reps {
				fmt.Printf("\033[2A")
			}
		}

		ui.PrintComplete()
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