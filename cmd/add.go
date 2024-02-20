package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add domains to a list of domains to be blocked",
	Long:  `Add domains to a list of domains to be blocked`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		domains := viper.GetStringSlice("default-domains")

		for _, domain := range args {
			if slices.Contains(domains, domain) {
				continue
			}

			domains = append(domains, domain)
		}

		viper.Set("default-domains", domains)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Domains added to the list")
	},
}

func init() {
	defaultsCmd.AddCommand(addCmd)
}
