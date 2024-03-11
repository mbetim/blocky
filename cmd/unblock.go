package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/mbetim/blocky/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock the domains for a x period of time",
	Long:  `Unblock the domains for a x period of time`,
	Run: func(cmd *cobra.Command, args []string) {
		domains := args

		if len(domains) == 0 {
			domains = viper.GetStringSlice("default-domains")
		}

		err := utils.UnblockDomains(domains)
		if err != nil {
			fmt.Println(err)
			return
		}

		timeToBlockAgain, _ := cmd.Flags().GetInt("time")
		joinedDomains := strings.Join(domains, "\n\t- ")
		messageFormat := "The following domains have been unblocked: \n\t- %s \n\nThey'll be blocked again in %d minutes or if any key is pressed\n"

		fmt.Printf(messageFormat, joinedDomains, timeToBlockAgain)

		time.AfterFunc(time.Duration(timeToBlockAgain)*time.Minute, blockDefaultDomains)

		var input string
		fmt.Scanln(&input)

		utils.BlockDomains(domains)
	},
}

func blockDefaultDomains() {
	fmt.Println("Blocking domains again")

	domains := viper.GetStringSlice("default-domains")
	err := utils.BlockDomains(domains)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.AddCommand(unblockCmd)

	unblockCmd.Flags().IntP("time", "t", 5, "Time, in minutes, to unblock the domains")
}
