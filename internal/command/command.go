package command

import (
	"fmt"
	"kegel-cli/internal/ui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: 		"kegel-trainer",
	Short:	"Kegel Trainer CLI is a tool for training one's pelvic floor muscles on the CLI.",
	Long:		"A simple tool that provides many training programs for men and women, who may be sitting too long on their desk jobs. Provides a simple an easy way to train their pelvic floor muscles anywhere.",
	RunE:		func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()

		choice, err := ui.ShowMenu()
		if err != nil {
			return nil
		}

		switch choice {
		case ui.MenuTrain:
			return startCmd.RunE(cmd, args)
		case ui.MenuProgress:
			fmt.Println("\n Coming soon!")
		case ui.MenuInfo:
			return infoCmd.RunE(cmd, args)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}