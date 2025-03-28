package cmd

import (
	"fmt"

	"github.com/mbetim/blocky/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block a domain",
	Long:  `Block a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		domains := args
		shouldBlockDefaultDomains, _ := cmd.Flags().GetBool("defaults")

		if shouldBlockDefaultDomains {
			domains = viper.GetStringSlice("default-domains")
		}

		if len(domains) == 0 {
			fmt.Println("requires at least 1 domain or use the flag --defaults")
			return
		}

		err := utils.BlockDomains(domains)
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.FlushDns()

		fmt.Println("The domains have been blocked")
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	blockCmd.Flags().BoolP("defaults", "d", false, "Block the domains in the `default-domains` list")
}
