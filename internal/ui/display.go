package ui

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	gray = color.New(color.FgHiBlack)
)

const barWidth = 30

func RenderBar(phase string, elapsed, total float64) {
	pct := math.Min(elapsed/total, 1.0)
	filled := int(math.Round(pct * float64(barWidth)))
	empty := barWidth - filled

	bar := buildBar(filled, empty, phase)
	remaining := total - elapsed
	if remaining < 0 {
		remaining = 0
	}

	fmt.Printf("\r\033[K%s  %s  %.0fs", phaseLabel(phase), bar, remaining)
}

func buildBar(filled, empty int, phase string) string {
	filledBlock := strings.Repeat("█", filled)
	emptyBlock := gray.Sprint(strings.Repeat("░", empty))

	var coloredFill string
	switch phase {
	case "squeeze":
		coloredFill = yellow.Sprint(filledBlock)
	case "rest":
		coloredFill = green.Sprint(filledBlock)

	}

	return fmt.Sprintf("[%s%s]", coloredFill, emptyBlock)
}

func phaseLabel(phase string) string {
	switch phase {
	case "squeeze":
		return yellow.Sprint("SQUEEZE")
	case "rest":
		return green.Sprint("REST   ")
	}
	return phase
}

func PrintBanner() {
	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Printf("\n")
	cyan.Println("  ███████╗██╗     ██╗        ██╗  ██╗███████╗ ██████╗ ███████╗██╗")
  cyan.Println("  ██╔════╝██║     ██║        ██║ ██╔╝██╔════╝██╔════╝ ██╔════╝██║")
  cyan.Println("  ██║     ██║     ██║ █████╗ █████╔╝ █████╗  ██║  ███╗█████╗  ██║")
  cyan.Println("  ██║     ██║     ██║ ╚════╝ ██╔═██╗ ██╔══╝  ██║   ██║██╔══╝  ██║")
  cyan.Println("  ███████╗███████╗██║        ██║  ██╗███████╗╚██████╔╝███████╗███████╗")
  cyan.Println("  ╚══════╝╚══════╝╚═╝        ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚══════╝╚══════╝")
  fmt.Println()
  cyan.Println("  ████████╗██████╗  █████╗ ██╗███╗   ██╗███████╗██████╗")
  cyan.Println("     ██╔══╝██╔══██╗██╔══██╗██║████╗  ██║██╔════╝██╔══██╗")
  cyan.Println("     ██║   ██████╔╝███████║██║██╔██╗ ██║█████╗  ██████╔╝")
  cyan.Println("     ██║   ██╔══██╗██╔══██║██║██║╚██╗██║██╔══╝  ██╔══██╗")
  cyan.Println("     ██║   ██║  ██║██║  ██║██║██║ ╚████║███████╗██║  ██║")
  cyan.Println("     ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝")
	fmt.Println()
}

func PrintPlanSummary(reps int, squeeze, rest time.Duration) {
	gray := color.New(color.FgHiBlack)
	bold := color.New(color.Bold)

	bold.Printf("  Reps     ")
	gray.Printf("%d\n", reps)
	bold.Printf("  Squeeze  ")
	gray.Printf("%0f\n", squeeze.Seconds())
	bold.Printf("  Rest     ")
	gray.Printf("%0f\n", rest.Seconds())
	bold.Printf("  Total    ")
	gray.Printf("~%.0f seconds\n\n", float64(reps)*(squeeze.Seconds()+rest.Seconds()))
}

func WaitForEnter() {
	bold := color.New(color.Bold)
	bold.Println("  Press Enter to begin...")
}

func PrintComplete() {
	green := color.New(color.FgGreen, color.Bold)
	gray := color.New(color.FgHiBlack)
	fmt.Println()
	green.Println("  ✓  Workout complete!")
	gray.Println("  Consistency is everything - see you next time!")
	fmt.Println()
}