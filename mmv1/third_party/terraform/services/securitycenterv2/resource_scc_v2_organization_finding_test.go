package securitycenterv2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccSecurityCenterV2OrganizationFinding_basic(t *testing.T) {
	t.Parallel()

	orgId := envvar.GetTestOrgFromEnv(t)
	suffix := acctest.RandString(t, 10)

	context := map[string]interface{}{
		"org_id":              orgId,
		"random_suffix":       randomSuffix,
		"finding":          "findingId-"  + randomSuffix,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterFinding_sccFindingBasicExample(orgId, suffix, "My description", "findingId1234"),
			},
			{
				ResourceName:      "google_scc_v2_organization_source.custom_finding",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecurityCenterFinding_sccFindingBasicExample(orgId, suffix, description string, finding string) string {
	return fmt.Sprintf(`
resource "google_scc_v2_organization_source" "custom_source" {
  display_name = "Terraform generated test"
  organization = "%{org_id}"
  description  = "%s"
}

resource "google_scc_v2_organization_finding" "custom_finding" {
  organization = "%{org_id}"
  source = google_scc_v2_organization_source.sourceId
  name  = "%{finding}"
}
`, description)
}
