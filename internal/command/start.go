package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:		"start",
	Short:	"Start a Kegel training session",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting session...")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}