package cmd

import (
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use: "status",
	Run: func(cmd *cobra.Command, args []string) {
		if err := search.PrintStatus(); err != nil {
			handleError(err)
		}
	},
}

func init() {
	searchCmd.AddCommand(statusCmd)
}
