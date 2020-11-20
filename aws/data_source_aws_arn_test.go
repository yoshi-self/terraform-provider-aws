package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceAwsArn_basic(t *testing.T) {
	resourceName := "data.aws_arn.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsArnConfig(endpoints.EuWest1RegionID),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceAwsArn(resourceName),
					resource.TestCheckResourceAttr(resourceName, "partition", "aws"),
					resource.TestCheckResourceAttr(resourceName, "service", "rds"),
					resource.TestCheckResourceAttr(resourceName, "region", endpoints.EuWest1RegionID),
					resource.TestCheckResourceAttr(resourceName, "account", "123456789012"),
					resource.TestCheckResourceAttr(resourceName, "resource", "db:mysql-db"),
				),
			},
		},
	})
}

func testAccDataSourceAwsArn(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		return nil
	}
}

//lintignore:AWSAT005
func testAccDataSourceAwsArnConfig(region string) string {
	return fmt.Sprintf(`
data "aws_arn" "test" {
  arn = "arn:aws:rds:%s:123456789012:db:mysql-db"
}
`, region)
}
