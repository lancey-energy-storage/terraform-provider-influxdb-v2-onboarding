package influxdbv2_onboarding

import (
	"context"
	"fmt"
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
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_USERNAME", "administrator"),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_PASSWORD", "Administrator1."),
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	options := influxdb.WithUserAndPass(d.Get("username").(string), d.Get("password").(string))
	influx, error := influxdb.New(d.Get("url").(string), "", options)
	if error != nil {
		return nil, fmt.Errorf("invalid InfluxDBv2 URL: %s", error)
	}

	err := influx.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging server: %s", err)
	}

	return influx, nil
}
