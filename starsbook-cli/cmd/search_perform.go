package cmd

import (
	"github.com/spf13/cobra"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"
	"strings"
)

// performCmd represents the execute command
var performCmd = &cobra.Command{
	Use: "perform",
	Run: func(cmd *cobra.Command, args []string) {
		searchTerm := strings.Join(args, " ")
		if err := search.Perform(searchTerm); err != nil {
			handleError(err)
		}
	},
}

func init() {
	searchCmd.AddCommand(performCmd)
}
