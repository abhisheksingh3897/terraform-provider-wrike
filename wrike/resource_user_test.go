package wrike

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccItem_Basic(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("wrike_user.user1", "email", "abhisheksingh17@ece.iiitp.ac.in"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "wrike_user" "user1" {    
    email= "abhisheksingh17@ece.iiitp.ac.in"
}
`)
}
