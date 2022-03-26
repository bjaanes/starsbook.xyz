package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"
	"os"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [index]",
	Short: "Delete a search key by it's id (see 'search keys list')",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argument (id) is required")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			handleError(err)
		}

		if err := search.DeleteApiKey(int64(id)); err != nil {
			handleError(err)
		}
	},
}

func init() {
	keysCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
