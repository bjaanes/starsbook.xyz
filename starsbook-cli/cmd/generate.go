package cmd

import (
	"github.com/spf13/cobra"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/download"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/genfrontend"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/genprojectfiles"
)

var skipDownloadImgs bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Downloads and generates all the things",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := conf.GetConfig()
		if err != nil {
			handleError(err)
			return
		}

		if err := download.NFTsFromIPFS(conf); err != nil {
			handleError(err)
			return
		}

		if !skipDownloadImgs {
			if err := download.ImgsFromIPFS(conf); err != nil {
				handleError(err)
				return
			}
		}

		if err := genprojectfiles.ProjectFiles(conf); err != nil {
			handleError(err)
			return
		}

		if err := genfrontend.ProjectFiles(conf); err != nil {
			handleError(err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolVarP(&skipDownloadImgs, "skip-download-images", "", false, "Skip downloading images")
}
