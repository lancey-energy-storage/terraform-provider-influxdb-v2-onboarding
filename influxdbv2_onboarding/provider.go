package influxdbv2_onboarding

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/influxdata/influxdb-client-go"
)

func Provider() *schema.Provider {
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
	influx := influxdb2.NewClient(d.Get("url").(string), "")
	_, err := influx.Ready(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging server: %s", err)
	}
	return influx, nil
}
