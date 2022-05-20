package cmd

import (
	"github.com/spf13/cobra"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/compress"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/download"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/folders"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/genfrontend"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/genprojectfiles"
)

var skipDownloadImgs bool
var skipIPFS bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Downloads and generates all the things",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := conf.GetConfig()
		if err != nil {
			handleError(err)
		}

		for _, p := range conf.Projects {
			if err := folders.CreateProjectFolderStructure(p); err != nil {
				handleError(err)
			}
		}

		if !skipIPFS {
			if err := download.NFTs(conf); err != nil {
				handleError(err)
			}

			if !skipDownloadImgs {
				if err := download.Imgs(conf); err != nil {
					handleError(err)
				}

				if err := download.ProjectImgs(conf); err != nil {
					handleError(err)
				}
				if err := compress.Images(conf); err != nil {
                                    handleError(err)
                                }

			}
		}

		if err := genprojectfiles.ProjectFiles(conf); err != nil {
			handleError(err)
		}

		if err := genfrontend.ProjectFiles(conf); err != nil {
			handleError(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolVar(&skipDownloadImgs, "skip-download-images", false, "Skip downloading images")
	generateCmd.Flags().BoolVar(&skipIPFS, "skip-ipfs", false, "Skip downloading stuff from IPFS")
}
