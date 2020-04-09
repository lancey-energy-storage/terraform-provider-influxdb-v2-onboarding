package influxdbv2_onboarding

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
	"github.com/rs/xid"
)

func ResourceSetup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSetupCreate,
		Read:   resourceSetupRead,
		Delete: resourceSetupDelete,
		Update: resourceSetupUpdate,
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceSetupCreate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	err := resourceSetupRead(d, meta)
	if err != nil {
		return fmt.Errorf("error getting status of influxdbv2 instance: %v", err)
	}
	if d.Get("allowed").(bool) {
		result, err := influx.Setup(context.Background(), d.Get("bucket").(string), d.Get("org").(string), d.Get("retention_period").(int))
		if err != nil {
			return fmt.Errorf("error setup endpoint: %v", err)
		}

		d.Set("token", result.Auth.Token)
		d.Set("user_id", result.User.ID)
		d.Set("bucket_id", result.Bucket.ID)
		d.Set("org_id", result.Org.ID)
		d.Set("auth_id", result.Auth.ID)
		id := xid.New().String()
		d.SetId(id)
	}

	return nil
}

func resourceSetupRead(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	result, err := influx.GetSetup()
	if err != nil {
		return fmt.Errorf("unable to call influxdbv2 instance: %v", err)
	}
	d.Set("allowed", result.Allowed)
	return nil
}

func resourceSetupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSetupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
