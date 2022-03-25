package cmd

import (
	"fmt"
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
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

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
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	generateCmd.Flags().BoolVarP(&skipDownloadImgs, "skip-download-images", "", false, "Skip downloading images")
}
