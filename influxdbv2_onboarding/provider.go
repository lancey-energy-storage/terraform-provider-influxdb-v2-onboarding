package influxdbv2_onboarding

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/lancey-energy-storage/influxdb-client-go"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"influxdbv2-onboarding_setup": ResourceSetup(),
		},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_URL", "http://localhost:9999"),
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	influx, err := influxdb.New(d.Get("url").(string), "")
	if err != nil {
		return nil, err
	}
	return influx, nil
}
