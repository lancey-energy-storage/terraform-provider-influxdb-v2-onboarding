package influxdbv2_onboarding

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestResourceSetup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigSetup(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("influxdbv2-onboarding_setup.setup", "bucket_id"),
					resource.TestCheckResourceAttrSet("influxdbv2-onboarding_setup.setup", "user_id"),
					resource.TestCheckResourceAttrSet("influxdbv2-onboarding_setup.setup", "org_id"),
					resource.TestCheckResourceAttrSet("influxdbv2-onboarding_setup.setup", "auth_id"),
					resource.TestCheckResourceAttrSet("influxdbv2-onboarding_setup.setup", "token"),
				),
			},
		},
	})
}

func testAccConfigSetup() string {
	return `
resource "influxdbv2-onboarding_setup" "setup" {
  username = "test"
  password = "test1234"
  bucket = "test-bucket"
  org = "test-org"
  retention_period = 4
}
`
}
