package cmd

import (
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		if err := search.ListApiKeys(); err != nil {
			handleError(err)
		}
	},
}

func init() {
	keysCmd.AddCommand(listCmd)
}
