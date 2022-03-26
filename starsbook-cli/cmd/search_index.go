package cmd

import (
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/search"

	"github.com/spf13/cobra"
)

var force bool

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Creates typesense indexes for search",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := conf.GetConfig()
		if err != nil {
			handleError(err)
		}

		if err := search.UpsertIndexes(conf, force); err != nil {
			handleError(err)
		}
	},
}

func init() {
	searchCmd.AddCommand(indexCmd)

	indexCmd.Flags().BoolVar(&force, "force", false, "Delete and recreate all search collections and indexes")
}
