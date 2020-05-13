package influxdbv2_onboarding

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestResourceSetup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckSetupExists("influxdbv2-onboarding_setup.setup"),
				),
			},
		},
	})
}

func testAccCheckSetupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.Attributes["bucket_id"] == "" {
			return fmt.Errorf("no bucket id set")
		}
		if rs.Primary.Attributes["user_id"] == "" {
			return fmt.Errorf("no user id set")
		}
		if rs.Primary.Attributes["org_id"] == "" {
			return fmt.Errorf("no org id set")
		}
		if rs.Primary.Attributes["auth_id"] == "" {
			return fmt.Errorf("no auth id set")
		}
		if rs.Primary.Attributes["token"] == "" {
			return fmt.Errorf("no token set")
		}

		return nil
	}
}

var testAccConfig = `
resource "influxdbv2-onboarding_setup" "setup" {
  username = "test"
  password = "test1234"
  bucket = "test-bucket"
  org = "test-org"
  retention_period = 4
}
`
