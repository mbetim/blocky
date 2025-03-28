package utils

import (
	"os/exec"

	"github.com/txn2/txeh"
)

func BlockDomains(domains []string) error {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		return err
	}

	savedDomains := make(map[string]struct{})

	for _, domain := range domains {
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
		return err
	}

	return nil
}

func UnblockDomains(domains []string) error {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		return err
	}

	for _, domain := range domains {
		wwwDomain := "www." + domain

		hosts.RemoveHost(domain)
		hosts.RemoveHost(wwwDomain)
	}

	hosts.RemoveHosts(domains)
	err = hosts.Save()
	if err != nil {
		return err
	}

	return nil
}

func FlushDns() {
	exec.Command("dscacheutil", "-flushcache").Run()
}
