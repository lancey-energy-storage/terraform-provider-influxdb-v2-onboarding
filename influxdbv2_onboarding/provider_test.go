package influxdbv2_onboarding

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

var testAccProviders = map[string]*schema.Provider{
	"influxdbv2-onboarding": Provider(),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("INFLUXDB_V2_URL"); v == "" {
		t.Fatal("INFLUXDB_V2_URL must be set for acceptance tests")
	}
}
