package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/lancey-energy-storage/terraform-provider-influxdb-v2-onboarding/influxdbv2_onboarding"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: influxdbv2_onboarding.Provider})
}
