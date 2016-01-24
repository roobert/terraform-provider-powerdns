package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/roobert/terraform-provider-powerdns/powerdns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: powerdns.Provider,
	})
}
