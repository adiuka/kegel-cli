package command

import (
	"fmt"
	"time"

	"kegel-cli/internal/progress"
	"kegel-cli/internal/ui"
	"kegel-cli/internal/workout"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:		"start",
	Short:	"Start a Kegel training session",
	RunE:	func(cmd *cobra.Command, args []string) error {
		return runWorkout()
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

func runWorkout() error {
	store, err := progress.Load()
	if err != nil {
		return err
	}

	plan := workout.ForLevel(store.CurrentLevel)

	ui.PrintPlanSummary(plan.Reps, plan.Squeeze, plan.Rest, plan.Name, store.SessionsAtCurrentLevel(), 10)
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

	store.AddSession(store.CurrentLevel, plan.Reps, int(plan.Squeeze.Seconds()), int(plan.Rest.Seconds()))

	if store.ShouldAdvance(10) {
		store.CurrentLevel ++
		ui.PrintLevelUp(workout.ForLevel(store.CurrentLevel))
	}

	return progress.Save(store)
}

func init() {
	rootCmd.AddCommand(startCmd)
}