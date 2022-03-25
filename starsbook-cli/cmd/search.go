package cmd

import (
	"github.com/spf13/cobra"
)

var searchTerm string

var searchCmd = &cobra.Command{
	Use: "search",
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
