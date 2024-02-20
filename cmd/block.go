package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/txn2/txeh"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block a domain",
	Long:  `Block a domain`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}

		savedDomains := make(map[string]struct{})

		for _, domain := range args {
			if _, ok := savedDomains[domain]; ok {
				continue
			}

			wwwDomain := "www." + domain

			hosts.AddHost("127.0.0.1", domain)
			hosts.AddHost("127.0.0.1", wwwDomain)

			savedDomains[domain] = struct{}{}
			savedDomains[wwwDomain] = struct{}{}
		}

		err = hosts.Save()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("The domains have been blocked")
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
