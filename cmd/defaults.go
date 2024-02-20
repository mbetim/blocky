package cmd

import (
	"github.com/spf13/cobra"
)

var defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Manage the list of default domains to block",
	Long:  `Manage the list of default domains to block`,
}

func init() {
	rootCmd.AddCommand(defaultsCmd)
}
