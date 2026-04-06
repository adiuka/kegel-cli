package command

import (
	"kegel-cli/internal/ui"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:		"info",
	Short:	"Learn about kegel training",
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintDescription()
		return nil
	},
}