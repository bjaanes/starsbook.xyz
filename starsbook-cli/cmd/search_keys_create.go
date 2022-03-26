package cmd

import (
	"fmt"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create \"description\"",
	Short: "Create a new search-only API key",
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		if description == "" {
			fmt.Println("Description is required")
			os.Exit(1)
		}

		if err := search.CreateApiKey(description); err != nil {
			handleError(err)
		}
	},
}

func init() {
	keysCmd.AddCommand(createCmd)
}
