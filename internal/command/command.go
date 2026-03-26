package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "kegel-trainer",
	Short: "Kegel Trainer CLI is a tool for training one's pelvic floor muscles on the CLI.",
	Long: "A simple tool that provides many training programs for men and women, who may be sitting too long on their desk jobs. Provides a simple an easy way to train their pelvic floor muscles anywhere.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}