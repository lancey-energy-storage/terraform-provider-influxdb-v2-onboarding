package influxdbv2_onboarding

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/influxdata/influxdb-client-go"
	"net/http"
)

func ResourceSetup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSetupCreate,
		Read:   resourceSetupRead,
		Delete: resourceSetupDelete,
		Update: resourceSetupUpdate,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_USERNAME", "administrator"),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_PASSWORD", "Administrator1."),
			},
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
	influx := meta.(influxdb2.Client)
	err := resourceSetupRead(d, meta)
	if err != nil {
		return fmt.Errorf("error getting status of influxdbv2 instance: %v", err)
	}
	if d.Get("allowed").(bool) {
		result, err := influx.Setup(context.Background(), d.Get("username").(string), d.Get("password").(string), d.Get("bucket").(string), d.Get("org").(string), d.Get("retention_period").(int))
		if err != nil {
			return fmt.Errorf("error setup endpoint: %v", err)
		}

		d.Set("token", result.Auth.Token)
		d.Set("user_id", result.User.Id)
		d.Set("bucket_id", result.Bucket.Id)
		d.Set("org_id", result.Org.Id)
		d.Set("auth_id", result.Auth.Id)
		id := ""
		url := influx.ServerUrl()
		id = url
		d.SetId(id)
	}

	return nil
}

func resourceSetupRead(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)

	var url = influx.ServerURL() + "/api/v2/setup"
	result, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while calling %s", url)
	}
	defer result.Body.Close()

	var jsonDecode map[string]interface{}
	json.NewDecoder(result.Body).Decode(&jsonDecode)
	d.Set("allowed", jsonDecode["allowed"])
	return nil
}

func resourceSetupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSetupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
